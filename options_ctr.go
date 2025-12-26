package choice4go

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/valyala/bytebufferpool"
)

type CTROptions[T Option] struct {
	kwOptions[T]
}

func (opt CTROptions[T]) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("CtrOptions{")
	opt.format(buff)
	buff.WriteString("}")

	return buff.String()
}

type CTRFinacialOptions struct {
	CTROptions[CTRFinacialOptions]

	code       string
	reportType ReportType
	reportName ReportName
}

func (opt CTRFinacialOptions) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("CtrOptions{")
	if opt.reportName != 0 {
		fmt.Fprintf(buff, " ReportName:%s ", opt.reportName)
		fmt.Fprintf(buff, "SecurityCode:%s ", opt.code)
	} else {
		fmt.Fprintf(buff, " SecurityCode:%s ", opt.code)
	}
	fmt.Fprintf(buff, "ReportType:%s", opt.reportType)

	opt.format(buff, "ReportName", "SecurityCode", "ReportType")

	buff.WriteString("}")

	return buff.String()
}

func (opt *CTRFinacialOptions) Code(v string) *CTRFinacialOptions {
	if v != "" {
		codeStr := fmt.Sprintf("SecuCode=%s", v)
		if optIdx := opt.findOptIdx("SecuCode"); optIdx < 0 {
			opt.options = append(opt.options, codeStr)
		} else {
			opt.options[optIdx] = codeStr
		}
		opt.code = v
	} else {
		slog.Error(
			"invalid security code",
			slog.String("code", v),
		)
	}

	return opt
}

func (opt *CTRFinacialOptions) Type(v ReportType) *CTRFinacialOptions {
	typeStr := fmt.Sprintf("ReportType=%d", v)

	if optIdx := opt.findOptIdx("ReportType"); optIdx < 0 {
		opt.options = append(opt.options, typeStr)
	} else {
		opt.options[optIdx] = typeStr
	}

	opt.reportType = v

	return opt
}

func (opt *CTRFinacialOptions) SetOption(key string, value any) (*CTRFinacialOptions, error) {
	if value == nil {
		slog.Warn(
			"ctr option set an nil value",
			slog.String("key", key),
		)
		return opt, nil
	}

	switch strings.ToLower(key) {
	case "reporttype", "report_type":
		switch v := value.(type) {
		case int:
			return opt.Type(ReportType(v)), nil
		case uint:
			return opt.Type(ReportType(v)), nil
		case uint8:
			return opt.Type(ReportType(v)), nil
		case int8:
			return opt.Type(ReportType(v)), nil
		case uint16:
			return opt.Type(ReportType(v)), nil
		case int16:
			return opt.Type(ReportType(v)), nil
		case uint32:
			return opt.Type(ReportType(v)), nil
		case int32:
			return opt.Type(ReportType(v)), nil
		case uint64:
			return opt.Type(ReportType(v)), nil
		case int64:
			return opt.Type(ReportType(v)), nil
		case string:
			if t, err := strconv.Atoi(v); err != nil {
				slog.Error(
					"ctr set option failed",
					slog.Any("error", err),
					slog.String("key", key),
					slog.Any("value", value),
				)
				return opt, err
			} else {
				return opt.Type(ReportType(t)), nil
			}
		default:
			slog.Error(
				"ctr set option unsupported ReportType value",
				slog.String("key", key),
				slog.Any("value", value),
			)
			return opt, ErrInvalidOptionValue
		}
	case "secucode", "secu_code", "symbol":
		if v, ok := value.(string); !ok {
			slog.Error(
				"ctr set option unsupported SecuCode value",
				slog.String("key", key),
				slog.Any("value", value),
			)
			return opt, ErrInvalidOptionValue
		} else {
			return opt.Code(v), nil
		}
	case "reportname", "report_name":
		if v, ok := value.(string); !ok {
			slog.Error(
				"ctr set option unsupported SecuCode value",
				slog.String("key", key),
				slog.Any("value", value),
			)
			return opt, ErrInvalidOptionValue
		} else {
			if err := opt.reportName.UnmarshalText([]byte(v)); err != nil {
				slog.Error(
					"ctr options parse report name failed",
					slog.Any("error", err),
				)

				return opt, errors.Join(ErrInvalidOptionValue, err)
			}
			return opt, nil
		}
	}

	if _, err := opt.kwOptions.SetOption(key, value); err != nil {
		return nil, err
	}

	return opt, nil
}

func GetReportName(opt Option) ReportName {
	ctr, ok := opt.(*CTRFinacialOptions)
	if !ok {
		return 0
	}

	return ctr.reportName
}

func GetReportDate(opt Option) time.Time {
	ctr, ok := opt.(*CTRFinacialOptions)
	if !ok {
		return time.Time{}
	}

	return ctr.reportDate
}

func GetReportType(opt Option) ReportType {
	ctr, ok := opt.(*CTRFinacialOptions)
	if !ok {
		return RptTypeCombined
	}

	return ctr.reportType
}

func NewCtrFinacialOptions() *CTRFinacialOptions {
	opt := new(CTRFinacialOptions)

	opt.initBase(opt.Type(RptTypeCombined))

	return opt
}

func NewCtrOptions[T Option]() *CTROptions[T] {
	opt := new(CTROptions[T])

	opt.initBase(opt.specificOpt)

	return opt
}
