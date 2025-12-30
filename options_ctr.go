package choice4go

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/valyala/bytebufferpool"
)

type CTROptions struct {
	kwOptions[CTROptions]

	code       string
	reportType ReportType
	reportName ReportName
}

func (opt CTROptions) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("CtrOptions{")
	if opt.reportName.reportName != RptNameInvalid {
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

// Code 设置报表标的代码
func (opt *CTROptions) Code(v string) *CTROptions {
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

func (opt CTROptions) GetCode() string { return opt.code }

// Type 设置报表类型
func (opt *CTROptions) Type(v ReportType) *CTROptions {
	typeStr := fmt.Sprintf("ReportType=%d", v)

	if optIdx := opt.findOptIdx("ReportType"); optIdx < 0 {
		opt.options = append(opt.options, typeStr)
	} else {
		opt.options[optIdx] = typeStr
	}

	opt.reportType = v

	return opt
}

func (opt CTROptions) GetType() ReportType { return opt.reportType }

func (opt *CTROptions) Name(v reportName) {
	opt.reportName.reportName = v
	opt.reportName.value = v.String()
}

func (opt CTROptions) GetName() ReportName { return opt.reportName }

func (opt *CTROptions) SetOption(key string, value any) (*CTROptions, error) {
	if value == nil {
		slog.Warn(
			"ctr option set an nil value",
			slog.String("key", key),
		)
		return opt, nil
	}

	switch strings.ToLower(strings.TrimSpace(key)) {
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
	case "secucode", "secu_code", "symbol", "security":
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
	case "name", "reportname", "report_name":
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

	return opt.kwOptions.SetOption(key, value)
}

// func GetReportName(opt Option) ReportName {
// 	ctr, ok := opt.(*CTROptions)
// 	if !ok {
// 		return ReportName{}
// 	}

// 	return ctr.reportName
// }

// func GetReportDate(opt Option) time.Time {
// 	ctr, ok := opt.(*CTROptions)
// 	if !ok {
// 		return time.Time{}
// 	}

// 	return ctr.reportDate
// }

// func GetReportType(opt Option) ReportType {
// 	ctr, ok := opt.(*CTROptions)
// 	if !ok {
// 		return RptTypeCombined
// 	}

// 	return ctr.reportType
// }

func NewCtrOptions() *CTROptions {
	opt := new(CTROptions)

	opt.initBase(opt.Type(RptTypeCombined))

	return opt
}
