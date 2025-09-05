package choice4go

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/../include/choice

#include "cgoChoiceBridge.h"
*/
import "C"
import (
	"log/slog"
)

//export cgoDataCallback
func cgoDataCallback(msg *C.EQMSG, _ C.LPVOID) C.int {
	if msg == nil {
		return 0
	}

	version := int(msg.version)

	switch msg.msgType {
	case C.eMT_err:
		slog.Error(
			"choice async query failed",
			slog.Any("err", singleton.Load().checkError(msg.err)),
			slog.Int("request_id", int(msg.requestID)),
			slog.Int("serial_id", int(msg.serialID)),
		)

		if msg.err == C.EQERR_LOGIN_DISCONNECT {
			if ins := singleton.Load(); ins != nil {
				if err := ins.Restart(); err != nil {
					slog.Error(
						"restart Choice failed",
						slog.Any("error", err),
					)
				}
			} else {
				slog.Error(
					"no singleton Choice instance to restart",
				)
			}
		}
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
