//go:build windows

package choice4go

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/../include/choice

#include "cgoChoiceBridge.h"
*/
import "C"
import (
	"io"
	"log/slog"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	cnDecoder = simplifiedchinese.GB18030.NewDecoder()
)

//export cgoLogCallback
func cgoLogCallback(log *C.char) C.int {
	rawMsg := C.GoString(log)
	rd := transform.NewReader(
		strings.NewReader(strings.TrimSpace(rawMsg)),
		cnDecoder,
	)

	if msg, err := io.ReadAll(rd); err != nil {
		slog.Error(
			"decode choide log callback msg failed",
			slog.Any("error", err),
			slog.String("raw_msg", rawMsg),
		)
	} else {
		slog.Debug(
			"choice log callback",
			slog.String("msg", string(msg)),
		)
	}

	return 0
}
