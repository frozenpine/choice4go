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
	started    atomic.Bool
	stopOnce   sync.Once

	rootCtx    context.Context
	rootCancel context.CancelFunc

	errMsgFn       C.err_getter
	dataReleaserFn C.data_releaser
	startFn        C.starter
	stopFn         C.stopper
	csdFn          C.query_pchar5_pdata
	cssFn          C.query_pchar3_pdata
	csecFn         C.query_pchar3_pdata
	tradedateFn    C.query_pchar3_pdata
	sectorFn       C.query_pchar3_pdata
	ctrFn          C.query_pchar3_pctrdata
	edbFn          C.query_pchar2_pdata
	edbDtlFn       C.query_pchar3_pdata
	cfnFn          C.query_cfn_pdata
	cfnDtlFn       C.query_pchar_pdata
}

func loadFuncErr() error {
	msg := C.dlerror()

	return fmt.Errorf(
		"%w: %s", ErrLoadLib, C.GoString(msg),
	)
}

func NewChoice(libDir, libName, cfgPath string) (ins *Choice, err error) {
	if ins = singleton.Load(); ins != nil {
		return
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

	ins = &Choice{
		libPath: filepath.Join(libDir, libName),
		cfgPath: cfgPath,
	}

	ins.loadOnce.Do(func() {
		cLibPath := C.CString(ins.libPath)
		defer C.free(unsafe.Pointer(cLibPath))

		if ins.lib = C.dlopen(cLibPath, C.RTLD_LAZY); ins.lib == nil {
			err = loadFuncErr()
			return
		}

		if fn := C.dlsym(ins.lib, C.ERR_GETTER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.errMsgFn = (C.err_getter)(fn)
		}

		if fn := C.dlsym(ins.lib, C.DATA_RELEASER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.dataReleaserFn = (C.data_releaser)(fn)
		}

		if fn := C.dlsym(ins.lib, C.CALLBACK_SETTER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			fnPtr := (C.callback_setter)(fn)

			if err = ins.checkError(C.CallCbSetter(
				fnPtr,
				C.datacallback(
					unsafe.Pointer(C.cDataCallback),
				),
			)); err != nil {
				return
			}
		}

		if fn := C.dlsym(ins.lib, C.STARTER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.startFn = (C.starter)(fn)
		}

		if fn := C.dlsym(ins.lib, C.STOPPER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.stopFn = (C.stopper)(fn)
		}

		if fn := C.dlsym(ins.lib, C.CSD_QUERIER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.csdFn = (C.query_pchar5_pdata)(fn)
		}

		if fn := C.dlsym(ins.lib, C.CSS_QUERIER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.cssFn = (C.query_pchar3_pdata)(fn)
		}

		if fn := C.dlsym(ins.lib, C.CSES_QUERIER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.csecFn = (C.query_pchar3_pdata)(fn)
		}

		if fn := C.dlsym(ins.lib, C.TRADEDATE_QUERIER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.tradedateFn = (C.query_pchar3_pdata)(fn)
		}

		if fn := C.dlsym(ins.lib, C.SECTOR_QUERIER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.sectorFn = (C.query_pchar3_pdata)(fn)
		}

		if fn := C.dlsym(ins.lib, C.CTR_QUERIER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.ctrFn = (C.query_pchar3_pctrdata)(fn)
		}

		if fn := C.dlsym(ins.lib, C.EDB_QUERIER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.edbFn = (C.query_pchar2_pdata)(fn)
		}

		if fn := C.dlsym(ins.lib, C.EDB_DTL_QUERIER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.edbDtlFn = (C.query_pchar3_pdata)(fn)
		}

		if fn := C.dlsym(ins.lib, C.CFN_QUERIER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.cfnFn = (C.query_cfn_pdata)(fn)
		}

		if fn := C.dlsym(ins.lib, C.CFN_DTL_QUERIER_NAME); fn == nil {
			err = loadFuncErr()
			return
		} else {
			ins.cfnDtlFn = (C.query_pchar_pdata)(fn)
		}
	})

	runtime.SetFinalizer(ins, func(ins *Choice) {
		ins.rootCancel()

		ins.unloadOnce.Do(func() {
			// 清理全局单例指针
			singleton.CompareAndSwap(ins, nil)

			if err := ins.Stop(); err != nil {
				slog.Error(
					"choice clean up stop failed",
					slog.Any("error", err),
				)
			}

			C.dlclose(ins.lib)

			ins.errMsgFn = nil
			ins.dataReleaserFn = nil
			ins.startFn = nil
			ins.stopFn = nil
			ins.csdFn = nil
			ins.cssFn = nil
			ins.csecFn = nil
			ins.tradedateFn = nil
			ins.sectorFn = nil
			ins.ctrFn = nil
			ins.edbFn = nil
			ins.edbDtlFn = nil
			ins.cfnFn = nil
			ins.cfnDtlFn = nil
		})
	})

	if !singleton.CompareAndSwap(nil, ins) {
		ins = singleton.Load()
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

func (ins *Choice) checkLibFn(name string) (fn *[0]byte, err error) {
	if ins.lib == nil {
		return nil, fmt.Errorf(
			"%w: lib not loaded", ErrInitialized,
		)
	}

	switch name {
	case "start":
		fn = ins.startFn
	case "stop":
		fn = ins.stopFn
	case "css":
		fn = ins.cssFn
	case "csd":
		fn = ins.csdFn
	case "cses":
		fn = ins.csecFn
	case "tradedates":
		fn = ins.tradedateFn
	case "sector":
		fn = ins.sectorFn
	case "ctr":
		fn = ins.ctrFn
	case "edb":
		fn = ins.edbFn
	case "edbquery":
		fn = ins.edbDtlFn
	case "cfn":
		fn = ins.cfnFn
	case "cfnquery":
		fn = ins.cfnDtlFn
	default:
		err = fmt.Errorf(
			"%w: unkown data function call %s", ErrLoadFunc, name,
		)
	}

	return
}

func (ins *Choice) releaseData(data *C.EQDATA) error {
	return ins.checkError(C.CallDataReleaser(
		ins.dataReleaserFn, unsafe.Pointer(data),
	))
}

func (ins *Choice) Start(
	ctx context.Context,
	user, pass string,
	options Option,
) (err error) {
	fn, err := ins.checkLibFn("start")
	if err != nil {
		return err
	}

	if ctx == nil {
		ctx = context.Background()
	}

	ins.startOnce.Do(func() {
		ins.rootCtx, ins.rootCancel = context.WithCancel(ctx)

		cUser := C.CString(user)
		cPass := C.CString(pass)

		var cOptions *C.char
		if options != nil {
			cOptions = C.CString(options.OptionString())
		}

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

		slog.Info(
			"choice run start with options",
			slog.String("user", user),
			slog.Any("options", options),
		)
		err = ins.checkError(C.CallStarter(
			fn, &login, cOptions,
			C.logcallback(unsafe.Pointer(C.cLogCallback)),
		))

		ins.started.Store(err == nil)
	})

	return
}

func (ins *Choice) Stop() (err error) {
	fn, err := ins.checkLibFn("stop")
	if err != nil {
		return err
	}

	if !ins.started.Load() {
		slog.Warn("choice api not started")
		return
	}

	ins.stopOnce.Do(func() {
		ins.rootCancel()

		err = ins.checkError(C.CallStopper(fn))
	})

	ins.started.Store(false)

	return
}

func convertStringArr(arr C.EQCHARARRAY) []string {
	if arr.nSize == 0 || arr.pChArray == nil {
		return nil
	}

	strArr := *(*[]C.EQCHAR)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(arr.pChArray)),
		Len:  int(arr.nSize),
		Cap:  int(arr.nSize),
	}))

	results := make([]string, 0, arr.nSize)
	for _, str := range strArr {
		var msg string
		if str.pChar != nil {
			msg = C.GoStringN(str.pChar, C.int(str.nSize)-1)
		} else {
			msg = ""
		}
		results = append(results, msg)
	}

	return results
}

func newEQValue(v *C.EQVARIENT) (*EQValue, error) {
	if v == nil {
		return nil, fmt.Errorf(
			"%w: empty EQVARIENT pointer", ErrDataEmpty,
		)
	}

	value, ok := valuePool.Get().(*EQValue)

	if !ok {
		return nil, fmt.Errorf(
			"%w: fail to get EQValue", ErrGetData,
		)
	}

	switch v.vtype {
	case C.eVT_null:
		value.valueType = ValueNull
	case C.eVT_char:
		value.valueType = ValueChar
	case C.eVT_byte:
		value.valueType = ValueByte
	case C.eVT_bool:
		value.valueType = ValueBool
	case C.eVT_short:
		value.valueType = ValueShort
	case C.eVT_ushort:
		value.valueType = ValueUShort
	case C.eVT_int:
		value.valueType = ValueInt
	case C.eVT_uInt:
		value.valueType = ValueUInt
	case C.eVT_int64:
		value.valueType = ValueInt64
	case C.eVT_uInt64:
		value.valueType = ValueUInt64
	case C.eVT_float:
		value.valueType = ValueSingle
	case C.eVT_double:
		value.valueType = ValueDouble
	case C.eVT_byteArray:
		value.valueType = ValueBytes
	case C.eVT_asciiString:
		value.valueType = ValueString
	case C.eVT_unicodeString:
		value.valueType = ValueString
	default:
		slog.Warn(
			"choice unsupported data value",
			slog.Any("type", v.vtype),
		)
	}

	if value.valueType == ValueString {
		value.valueString = C.GoStringN(
			v.eqchar.pChar, C.int(v.eqchar.nSize)-1,
		)
		clear(value.valueBuffer[:])
	} else {
		C.memcpy(
			unsafe.Pointer(&value.valueBuffer[0]),
			unsafe.Pointer(&v.unionValues[0]),
			8,
		)
		value.valueString = ""
	}

	return value, nil
}

func newEQData(v *C.EQDATA) (*EQData, error) {
	if v == nil || v.valueArray.nSize == 0 {
		return nil, ErrDataEmpty
	}

	if v.valueArray.nSize != v.codeArray.nSize*
		v.dateArray.nSize*v.indicatorArray.nSize {
		return nil, fmt.Errorf(
			"%w: value buffer[%d], date len[%d], indicator len[%d]",
			ErrDataLenMissMatch, v.valueArray.nSize,
			v.dateArray.nSize, v.indicatorArray.nSize,
		)
	}

	data, ok := dataPool.Get().(*EQData)
	if !ok {
		return nil, fmt.Errorf(
			"%w: fail to get EQData", ErrGetData,
		)
	}

	data.codes = convertStringArr(v.codeArray)
	data.indicators = convertStringArr(v.indicatorArray)
	data.dateList = convertStringArr(v.dateArray)
	data.values = make([]*EQValue, v.valueArray.nSize)

	values := *(*[]C.EQVARIENT)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(v.valueArray.pEQVarient)),
		Len:  int(v.valueArray.nSize),
		Cap:  int(v.valueArray.nSize),
	}))

	var err error
	for idx, v := range values {
		if data.values[idx], err = newEQValue(&v); err != nil {
			return nil, err
		}
	}

	runtime.SetFinalizer(data, func(data *EQData) {
		for _, value := range data.values {
			valuePool.Put(value)
		}

		dataPool.Put(data)
	})

	return data, nil
}

func newEQCtrData(v *C.EQCTRDATA) (*EQCtrData, error) {
	if v == nil || v.valueArray.nSize == 0 {
		return nil, ErrDataEmpty
	}

	if v.valueArray.nSize != C.uint(v.row*v.column) {
		return nil, fmt.Errorf(
			"%w: value buffer[%d], row[%d], column[%d]",
			ErrDataLenMissMatch, v.valueArray.nSize,
			v.row, v.column,
		)
	}

	data, ok := ctrDataPool.Get().(*EQCtrData)
	if !ok {
		return nil, fmt.Errorf(
			"%w: fail to get EQCtrData", ErrGetData,
		)
	}

	data.row = int(v.row)
	data.column = int(v.column)
	data.indicators = convertStringArr(v.indicatorArray)
	data.values = make([]*EQValue, v.valueArray.nSize)

	values := *(*[]C.EQVARIENT)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(v.valueArray.pEQVarient)),
		Len:  int(v.valueArray.nSize),
		Cap:  int(v.valueArray.nSize),
	}))

	var err error
	for idx, v := range values {
		if data.values[idx], err = newEQValue(&v); err != nil {
			return nil, err
		}
	}

	runtime.SetFinalizer(data, func(data *EQCtrData) {
		for _, value := range data.values {
			valuePool.Put(value)
		}

		ctrDataPool.Put(data)
	})

	return data, nil
}

func checkCommonArgs(
	codes, indicators []string, options Option,
) (*C.char, *C.char, *C.char, error) {
	if len(codes) <= 0 || len(indicators) <= 0 {
		return nil, nil, nil, fmt.Errorf(
			"%w: codes or indicators is empty", ErrInvalidArgs,
		)
	}

	if len(indicators) > MAX_INDICATOR_COUNT {
		return nil, nil, nil, fmt.Errorf(
			"%w: exceed indicator count, max %d",
			ErrInvalidArgs, MAX_INDICATOR_COUNT)
	}

	var cOptions *C.char
	if options != nil {
		cOptions = C.CString(options.OptionString())
	}

	return C.CString(strings.Join(codes, ",")),
		C.CString(strings.Join(indicators, ",")),
		cOptions, nil
}

func (ins *Choice) callPData(
	fn *[0]byte, args ...*C.char,
) (*EQData, error) {
	defer func() {
		for _, ptr := range args {
			C.free(unsafe.Pointer(ptr))
		}
	}()

	var (
		pData *C.EQDATA
		rtn   C.EQErr
	)

	switch len(args) {
	case 1:
		rtn = C.CallPCharPData(fn, args[0], &pData)
	case 2:
		rtn = C.CallPChar2PData(fn, args[0], args[1], &pData)
	case 3:
		rtn = C.CallPChar3PData(fn, args[0], args[1], args[2], &pData)
	case 5:
		rtn = C.CallPChar5PData(
			fn, args[0], args[1], args[2], args[3], args[4], &pData,
		)
	default:
		return nil, fmt.Errorf(
			"%w: unsupported args count: %d", ErrInvalidArgs, len(args),
		)
	}

	if err := ins.checkError(rtn); err != nil {
		return nil, err
	}
	defer ins.releaseData(pData)

	return newEQData(pData)
}

func (ins *Choice) Csd(
	codes, indicators []string,
	start, end time.Time,
	options Option,
) (*EQData, error) {
	fn, err := ins.checkLibFn("csd")
	if err != nil {
		return nil, err
	}

	cCodes, cIndicators, cOptions, err := checkCommonArgs(codes, indicators, options)
	if err != nil {
		return nil, err
	}

	cStart := C.CString(start.Format("2006-01-02"))
	cEnd := C.CString(end.Format("2006-01-02"))

	return ins.callPData(
		fn, cCodes, cIndicators, cStart, cEnd, cOptions,
	)
}

func (ins *Choice) Css(
	codes, indicators []string, options Option,
) (*EQData, error) {
	fn, err := ins.checkLibFn("css")
	if err != nil {
		return nil, err
	}

	cCodes, cIndicators, cOptions, err := checkCommonArgs(codes, indicators, options)
	if err != nil {
		return nil, err
	}

	return ins.callPData(
		fn, cCodes, cIndicators, cOptions,
	)
}

func (ins *Choice) CSec(
	blockCodes, indicators []string, options Option,
) (*EQData, error) {
	fn, err := ins.checkLibFn("csec")
	if err != nil {
		return nil, err
	}

	if len(blockCodes) > 6 {
		return nil, fmt.Errorf(
			"%w: block code count exceeded 6", ErrInvalidArgs,
		)
	}

	cCodes, cIndicators, cOptions, err := checkCommonArgs(
		blockCodes, indicators, options,
	)
	if err != nil {
		return nil, err
	}

	return ins.callPData(
		fn, cCodes, cIndicators, cOptions,
	)
}

func (ins *Choice) TradeDates(
	start, end time.Time,
	options Option,
) (*EQData, error) {
	fn, err := ins.checkLibFn("tradedates")
	if err != nil {
		return nil, err
	}

	cStart := C.CString(start.Format("2006-01-02"))
	cEnd := C.CString(end.Format("2006-01-02"))

	var cOptions *C.char
	if options != nil {
		cOptions = C.CString(options.OptionString())
	}
	defer func() {
		C.free(unsafe.Pointer(cStart))
		C.free(unsafe.Pointer(cEnd))
		C.free(unsafe.Pointer(cOptions))
	}()

	return ins.callPData(fn, cStart, cEnd, cOptions)
}
