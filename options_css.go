package choice4go

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/valyala/bytebufferpool"
)

type CSSOptions struct {
	kwOptions[CSSOptions]

	ttm       ttmType
	tableName string
}

func (opt CSSOptions) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("CssOptions{")
	opt.format(buff)
	buff.WriteString("}")

	return buff.String()
}

// TTMType 设置TTM基准日
func (opt *CSSOptions) TTMType(ttm ttmType) *CSSOptions {
	typeStr := fmt.Sprintf("TTMType=%d", ttm)

	if optIdx := opt.findOptIdx("TTMType"); optIdx < 0 {
		opt.options = append(opt.options, typeStr)
	} else {
		opt.options[optIdx] = typeStr
	}

	opt.ttm = ttm

	return opt
}

func (opt CSSOptions) GetTTM() ttmType { return opt.ttm }

func (opt *CSSOptions) TableName(v string) (*CSSOptions, error) {
	opt.tableName = v
	return opt, nil
}

func (opt CSSOptions) GetTableName() string { return opt.tableName }

func (opt *CSSOptions) SetOption(key string, value any) (*CSSOptions, error) {
	if value == nil {
		slog.Warn(
			"css option set an nil value",
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
	case "ttm", "ttmtype", "ttm_type":
		switch v := value.(type) {
		case string:
			d, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			opt.ttm = ttmType(d)
		case []byte:
			d, err := strconv.Atoi(string(v))
			if err != nil {
				return nil, err
			}
			opt.ttm = ttmType(d)
		case int:
			opt.ttm = ttmType(v)
		case uint:
			opt.ttm = ttmType(v)
		case int8:
			opt.ttm = ttmType(v)
		case uint8:
			if v >= '0' && v <= '9' {
				opt.ttm = ttmType(v - '0')
			} else {
				opt.ttm = ttmType(v)
			}
		case int16:
			opt.ttm = ttmType(v)
		case uint16:
			opt.ttm = ttmType(v)
		case int32:
			opt.ttm = ttmType(v)
		case uint32:
			opt.ttm = ttmType(v)
		case int64:
			opt.ttm = ttmType(v)
		case uint64:
			opt.ttm = ttmType(v)
		default:
			slog.Error(
				"css set option unsupported ttm type value",
				slog.String("key", key),
				slog.Any("value", value),
			)
			return opt, ErrInvalidOptionValue
		}
	}

	return opt.kwOptions.SetOption(key, value)
}

func NewCssOptions() *CSSOptions {
	opt := new(CSSOptions)

	return opt.initBase(opt)
}
