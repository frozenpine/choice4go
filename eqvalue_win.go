//go:build windows

package choice4go

import (
	"io"
	"log/slog"
	"strings"

	"golang.org/x/text/transform"
)

func (v *EQValue) GetString() string {
	if decoded, err := io.ReadAll(
		transform.NewReader(
			strings.NewReader(v.valueString),
			cnDecoder,
		),
	); err != nil {
		slog.Error(
			"Decode GB18030 failed",
			slog.Any("error", err),
		)
		return v.valueString
	} else {
		return string(decoded)
	}
}
