package choice4go

import (
	"log/slog"
	"strings"

	"github.com/valyala/bytebufferpool"
)

type CSDOptions struct {
	kwOptions[CSDOptions]

	tableName string
}

func (opt CSDOptions) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("CsdOptions{")
	opt.format(buff)
	buff.WriteString("}")

	return buff.String()
}

func (opt *CSDOptions) TableName(v string) (*CSDOptions, error) {
	opt.tableName = v
	return opt, nil
}

func (opt CSDOptions) GetTableName() string { return opt.tableName }

func (opt *CSDOptions) SetOption(key string, value any) (*CSDOptions, error) {
	if value == nil {
		slog.Warn(
			"csd option set an nil value",
			slog.String("key", key),
		)
		return opt, nil
	}

	switch strings.ToLower(strings.TrimSpace(key)) {
	case "table", "name", "table_name", "tablename":
		switch v := value.(type) {
		case string:
			return opt.TableName(v)
		case []byte:
			return opt.TableName(string(v))
		default:
			slog.Error(
				"csd set option unsupported table name value",
				slog.String("key", key),
				slog.Any("value", value),
			)
			return opt, ErrInvalidOptionValue
		}
	}

	return opt.kwOptions.SetOption(key, value)
}

func NewCsdOptions() *CSDOptions {
	opt := new(CSDOptions)

	return opt.initBase(opt)
}
