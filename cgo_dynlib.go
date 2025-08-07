package choice4go

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/dependency/includes
#cgo LDFLAGS: -ldl

#include <string.h>

#include "cgoDynLib.h"

extern int cLogCallback(const char* pLog);
extern int cDataCallback(const EQMSG* pMsg, LPVOID lpUserParam);
*/
import "C"
import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	MAX_INDICATOR_COUNT = 64
)

var (
	singleton atomic.Pointer[Choice]
)

type Choice struct {
	libPath string
	cfgPath string
	lib     unsafe.Pointer

	loadOnce   sync.Once
	unloadOnce sync.Once
	startOnce  sync.Once
	stopOnce   sync.Once

	rootCtx    context.Context
	rootCancel context.CancelFunc

	errMsgFn       C.err_getter
	dataReleaserFn C.data_releaser
	startFn        C.starter
	stopFn         C.stopper
	csdFn          C.query_pchar5_pdata
}

func NewChoice(libDir, libName, cfgPath string) (instance *Choice, err error) {
	if ins := singleton.Load(); ins != nil {
		return ins, nil
	}

	libIdentity := strings.SplitN(libName, ".", 2)

	switch runtime.GOOS {
	case "linux":
		if !strings.HasPrefix(libIdentity[0], "lib") {
			libName = fmt.Sprintf("lib%s.so", libIdentity[0])
		} else {
			libName = fmt.Sprintf("%s.so", libIdentity[0])
		}

		if err := os.Setenv("LD_LIBRARY_PATH", libDir); err != nil {
			return nil, err
		}

		slog.Info(
			"lib environment setted for linux",
			slog.String("LD_LIBRARY_PATH", os.Getenv("LD_LIBRARY_PATH")),
		)
	case "windows":
		libName = fmt.Sprintf("%s.dll", libIdentity[0])
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedSys, runtime.GOOS)
	}

	instance = &Choice{
		libPath: filepath.Join(libDir, libName),
		cfgPath: cfgPath,
	}

	instance.loadOnce.Do(func() {
		cLibPath := C.CString(instance.libPath)
		defer C.free(unsafe.Pointer(cLibPath))

		instance.lib = C.dlopen(cLibPath, C.RTLD_LAZY)
		if instance.lib == nil {
			msg := C.dlerror()

			err = fmt.Errorf(
				"%w: %s", ErrLoadLib, C.GoString(msg),
			)
			return
		}

		if fn := C.dlsym(
			instance.lib,
			C.ERR_GETTER_NAME,
		); fn == nil {
			msg := C.dlerror()

			err = fmt.Errorf(
				"%w: %s", ErrLoadFunc, C.GoString(msg),
			)

			return
		} else {
			instance.errMsgFn = (C.err_getter)(fn)
		}

		if fn := C.dlsym(
			instance.lib,
			C.DATA_RELEASER_NAME,
		); fn == nil {
			msg := C.dlerror()

			err = fmt.Errorf(
				"%w: %s", ErrLoadFunc, C.GoString(msg),
			)
			return
		} else {
			instance.dataReleaserFn = (C.data_releaser)(fn)
		}

		if fn := C.dlsym(
			instance.lib,
			C.CALLBACK_SETTER_NAME,
		); fn == nil {
			msg := C.dlerror()

			err = fmt.Errorf(
				"%w: %s", ErrLoadFunc, C.GoString(msg),
			)
			return
		} else {
			fnPtr := (C.callback_setter)(fn)

			if err = instance.checkError(C.CallCbSetter(
				fnPtr,
				C.datacallback(
					unsafe.Pointer(C.cDataCallback),
				),
			)); err != nil {
				return
			}
		}

		if fn := C.dlsym(
			instance.lib,
			C.STARTER_NAME,
		); fn == nil {
			msg := C.dlerror()

			err = fmt.Errorf(
				"%w: %s", ErrLoadFunc, C.GoString(msg),
			)
			return
		} else {
			instance.startFn = (C.starter)(fn)
		}

		if fn := C.dlsym(
			instance.lib,
			C.STOPPER_NAME,
		); fn == nil {
			msg := C.dlerror()

			err = fmt.Errorf(
				"%w: %s", ErrLoadFunc, C.GoString(msg),
			)

			return
		} else {
			instance.stopFn = (C.stopper)(fn)
		}

		if fn := C.dlsym(
			instance.lib,
			C.CSD_QUERIER_NAME,
		); fn == nil {
			msg := C.dlerror()

			err = fmt.Errorf(
				"%w: %s", ErrLoadFunc, C.GoString(msg),
			)
			return
		} else {
			instance.csdFn = (C.query_pchar5_pdata)(fn)
		}
	})

	runtime.SetFinalizer(instance, func(ins *Choice) {
		ins.rootCancel()

		ins.unloadOnce.Do(func() {
			if err := ins.Stop(); err != nil {
				slog.Error(
					"choice clean up failed",
					slog.Any("error", err),
				)
			}

			C.dlclose(ins.lib)

			ins.startFn = nil
			ins.stopFn = nil
			ins.csdFn = nil
		})
	})

	if !singleton.CompareAndSwap(nil, instance) {
		instance = singleton.Load()
	}

	return
}

func (ins *Choice) getErrString(code C.EQErr) string {
	if ins.errMsgFn == nil {
		panic("no err msg getter found")
	}

	msg := C.CallErrGetter(
		ins.errMsgFn, code, C.eLang_en,
	)

	return C.GoString(msg)
}

func (ins *Choice) checkError(code C.EQErr) error {
	if code == 0 {
		return nil
	}

	msg := ins.getErrString(code)

	return fmt.Errorf("%w: [%d] %s", ErrEQCall, code, msg)
}

func (ins *Choice) checkLib() error {
	if ins.lib == nil {
		return fmt.Errorf(
			"%w: lib not loaded", ErrInitialized,
		)
	}

	return nil
}

func (ins *Choice) releaseData(data *C.EQDATA) error {
	return ins.checkError(C.CallDataReleaser(
		ins.dataReleaserFn, unsafe.Pointer(data),
	))
}

func (ins *Choice) Start(ctx context.Context, user, pass string) (err error) {
	if err = ins.checkLib(); err != nil {
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	ins.startOnce.Do(func() {
		ins.rootCtx, ins.rootCancel = context.WithCancel(ctx)

		cUser := C.CString(user)
		cPass := C.CString(pass)
		cOptions := C.CString("ForceLogin=1,LogLevel=1")

		defer func() {
			C.free(unsafe.Pointer(cUser))
			C.free(unsafe.Pointer(cPass))
			C.free(unsafe.Pointer(cOptions))
		}()

		login := C.EQLOGININFO{}
		C.memcpy(
			unsafe.Pointer(&login.userName[0]),
			unsafe.Pointer(cUser),
			255,
		)
		C.memcpy(
			unsafe.Pointer(&login.password[0]),
			unsafe.Pointer(cPass),
			255,
		)

		err = ins.checkError(C.CallStarter(
			ins.startFn,
			&login, cOptions,
			C.logcallback(unsafe.Pointer(C.cLogCallback)),
		))
	})

	return
}

func (ins *Choice) Stop() (err error) {
	if err = ins.checkLib(); err != nil {
		return err
	}

	ins.stopOnce.Do(func() {
		err = ins.checkError(C.CallStopper(ins.stopFn))
	})

	return
}

func convertStringArr(arr C.EQCHARARRAY) []string {
	if arr.nSize == 0 {
		return nil
	}

	strArr := *(*[]C.EQCHAR)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(arr.pChArray)),
		Len:  int(arr.nSize),
		Cap:  int(arr.nSize),
	}))

	results := make([]string, 0, arr.nSize)
	for _, str := range strArr {
		msg := C.GoStringN(str.pChar, C.int(str.nSize)-1)
		results = append(results, msg)
	}

	return results
}

func convertEQData(data *C.EQDATA) (*EQData, error) {
	if data.valueArray.nSize == 0 {
		return nil, ErrDataEmpty
	}

	if data.valueArray.nSize != data.dateArray.nSize*data.indicatorArray.nSize {
		return nil, fmt.Errorf(
			"%w: value buffer[%d], date len[%d], indicator len[%d]",
			ErrDataLenMissMatch, data.valueArray.nSize,
			data.dateArray.nSize, data.indicatorArray.nSize,
		)
	}

	result := EQData{
		Codes:      convertStringArr(data.codeArray),
		Indicators: convertStringArr(data.indicatorArray),
		Date:       convertStringArr(data.dateArray),
		Values:     make([]EQValue, data.valueArray.nSize),
	}

	values := *(*[]C.EQVARIENT)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&data.valueArray.pEQVarient)),
		Len:  int(data.valueArray.nSize),
		Cap:  int(data.valueArray.nSize),
	}))

	for idx, v := range values {
		switch v.vtype {
		case C.eVT_null:
			result.Values[idx].ValueType = ValueNull
		case C.eVT_char:
			result.Values[idx].ValueType = ValueChar
		case C.eVT_byte:
			result.Values[idx].ValueType = ValueByte
		case C.eVT_bool:
			result.Values[idx].ValueType = ValueBool
		case C.eVT_short:
			result.Values[idx].ValueType = ValueShort
		case C.eVT_ushort:
			result.Values[idx].ValueType = ValueUShort
		case C.eVT_int:
			result.Values[idx].ValueType = ValueInt
		case C.eVT_uInt:
			result.Values[idx].ValueType = ValueUInt
		case C.eVT_int64:
			result.Values[idx].ValueType = ValueInt64
		case C.eVT_uInt64:
			result.Values[idx].ValueType = ValueUInt64
		case C.eVT_float:
			result.Values[idx].ValueType = ValueSingle
		case C.eVT_double:
			result.Values[idx].ValueType = ValueDouble
		case C.eVT_byteArray:
			result.Values[idx].ValueType = ValueBytes
		case C.eVT_asciiString:
			result.Values[idx].ValueType = ValueString
		case C.eVT_unicodeString:
			result.Values[idx].ValueType = ValueString
		default:
			slog.Warn(
				"choice unsupported data value",
				slog.Any("type", v.vtype),
			)
		}

		C.memcpy(
			unsafe.Pointer(&result.Values[idx].valueBuffer[0]),
			unsafe.Pointer(&v.unionValues[0]),
			8,
		)
	}

	return &result, nil
}

func (ins *Choice) Csd(
	codes []string,
	indicators []string,
	start, end time.Time,
	options []string,
) (*EQData, error) {
	if err := ins.checkLib(); err != nil {
		return nil, err
	}

	if len(codes) <= 0 {
		return nil, fmt.Errorf(
			"%w: codes is empty", ErrInvalidCodes,
		)
	}

	if len(indicators) > MAX_INDICATOR_COUNT {
		return nil, fmt.Errorf("%w: max %d", ErrTooManyIndicators, MAX_INDICATOR_COUNT)
	}

	cCodes := C.CString(strings.Join(codes, ","))
	cIndicators := C.CString(strings.Join(indicators, ","))
	cStart := C.CString(start.Format("2006-01-02"))
	cEnd := C.CString(end.Format("2006-01-02"))

	var cOptions *C.char
	if len(options) > 0 {
		cOptions = C.CString(strings.Join(options, ","))
	}

	defer func() {
		C.free(unsafe.Pointer(cCodes))
		C.free(unsafe.Pointer(cIndicators))
		C.free(unsafe.Pointer(cStart))
		C.free(unsafe.Pointer(cEnd))
		C.free(unsafe.Pointer(cOptions))
	}()

	var data *C.EQDATA

	if err := ins.checkError(C.CallPChar5PData(
		ins.csdFn,
		cCodes, cIndicators,
		cStart, cEnd, cOptions,
		&data,
	)); err != nil {
		return nil, err
	}
	defer ins.releaseData(data)

	return convertEQData(data)
}
