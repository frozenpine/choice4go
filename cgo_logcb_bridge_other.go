//go:build !windows

package choice4go

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/../include/choice

#include "cgoChoiceBridge.h"
*/
import "C"

import "log/slog"

//export cgoLogCallback
func cgoLogCallback(log *C.char) C.int {
	slog.Debug(
		"choice log callback",
		slog.String("msg", C.GoString(log)),
	)

	return 0
}
