package choice4go

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/dependency/includes

#include "cgoChoiceBridge.h"
*/
import "C"
import (
	"log/slog"
)

//export cgoLogCallback
func cgoLogCallback(log *C.char) C.int {
	msg := C.GoString(log)

	slog.Debug(
		"choice log callback",
		slog.String("msg", msg),
	)

	return 0
}

//export cgoDataCallback
func cgoDataCallback(msg *C.EQMSG, _param C.LPVOID) C.int {
	version := int(msg.version)

	switch msg.msgType {
	case C.eMT_err:
		slog.Error(
			"choice async query failed",
			slog.Any("err", singleton.Load().checkError(msg.err)),
			slog.Int("request_id", int(msg.requestID)),
			slog.Int("serial_id", int(msg.serialID)),
		)
	case C.eMT_response:
		slog.Info("choice async query response")
	case C.eMT_partialResponse:
		slog.Info("choice async query partial response")
	case C.eMT_others:
		slog.Info("choice other info")
	default:
		slog.Error(
			"choice unkown msg type",
			slog.Int("version", version),
			slog.Any("msg_type", msg.msgType),
		)
	}

	return 0
}
