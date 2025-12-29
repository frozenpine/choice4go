package choice4go

import (
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strings"
	"time"
)

type Option interface {
	fmt.Stringer

	OptionString() string
	GetOption(string) string
	GetOptions(...string) []string
}

//go:generate stringer -type period -linecomment
type period uint8

const (
	Daily   period = 1 // 日频
	Weekly  period = 2 // 周频
	Monthly period = 3 // 月频
	Yearly  period = 4 // 年频
)

//go:generate stringer -type adjustFlag -linecomment
type adjustFlag uint8

const (
	NoAdjusted       adjustFlag = 1 // 不复权
	BackwordAdjusted adjustFlag = 2 // 后复权
	ForwardAdjusted  adjustFlag = 3 // 前复权
)

//go:generate stringer -type currency -linecomment
type currency uint8

const (
	CurrOrigin currency = 1 // 原始币种
	CurrCNY    currency = 2 // 人民币
	CurrUSD    currency = 3 // 美元
	CurrHKD    currency = 4 // 港元
)

//go:generate stringer -type bondType -linecomment
type bondType uint8

const (
	BondClean bondType = 1 // 净价
	BondDirty bondType = 2 // 全价
	BondROI   bondType = 3 // 收益率
)

//go:generate stringer -type sortOrder -linecomment
type sortOrder uint8

const (
	DateASC  sortOrder = 1
	DateDESC sortOrder = 2
)

type options []string

func (opt options) OptionString() string {
	return strings.Join(opt, ",")
}

func (opt options) findOptIdx(prefix ...string) int {
	return slices.IndexFunc(opt, func(v string) bool {
		for _, match := range prefix {
			if strings.HasPrefix(v, match) {
				return true
			}
		}

		return false
	})
}

func (opt options) GetOption(key string) string {
	if idx := opt.findOptIdx(key); idx < 0 {
		return ""
	} else {
		return strings.TrimPrefix(opt[idx], key+"=")
	}
}

func (opt options) GetOptions(keys ...string) []string {
	results := make([]string, len(keys))
	for idx, key := range keys {
		if opt := opt.GetOption(key); opt == "" {
			return nil
		} else {
			results[idx] = opt
		}
	}

	return results
}

type BaseOption[T Option] interface {
	Option

	Period(period) *T
	Adjust(adjustFlag) *T
	Currency(currency) *T
	BondType(bondType) *T
	DateASC() *T
	DateDESC() *T
	FillData() *T
}

type baseOptions[T Option] struct {
	options

	specificOpt *T
	period      period
	adjustFlag  adjustFlag
	currType    currency
	bondType    bondType
	dateDESC    bool
	fillFlag    bool
}

func (opt *baseOptions[T]) initBase(ptr *T) *T {
	opt.specificOpt = ptr
	opt.period = Daily
	opt.adjustFlag = NoAdjusted
	opt.currType = CurrOrigin
	opt.bondType = BondClean

	return opt.specificOpt
}

func (opt *baseOptions[T]) format(wr io.Writer) {
	fmt.Fprintf(wr, "Period:%+v ", opt.period)
	fmt.Fprintf(wr, "Adjust:%+v ", opt.adjustFlag)
	fmt.Fprintf(wr, "Currency:%+v ", opt.currType)
	fmt.Fprintf(wr, "BondPrice:%+v ", opt.bondType)
	if opt.dateDESC {
		fmt.Fprintf(wr, "DateSort:DESC")
	} else {
		fmt.Fprintf(wr, "DateSort:ASC")
	}
}

// Period 设置数据频次
func (opt *baseOptions[T]) Period(p period) *T {
	periodOpt := fmt.Sprintf("Period=%d", p)

	if optIdx := opt.findOptIdx("Period"); optIdx < 0 {
		opt.options = append(opt.options, periodOpt)
	} else {
		opt.options[optIdx] = periodOpt
	}

	opt.period = p

	return opt.specificOpt
}

func (opt *baseOptions[T]) GetPeriod() period { return opt.period }

// Adjust 设置复权方式
func (opt *baseOptions[T]) Adjust(flag adjustFlag) *T {
	flagOpt := fmt.Sprintf("AdjustFlag=%d", flag)

	if optIdx := opt.findOptIdx("AdjustFlag"); optIdx < 0 {
		opt.options = append(opt.options, flagOpt)
	} else {
		opt.options[optIdx] = flagOpt
	}

	opt.adjustFlag = flag
	return opt.specificOpt
}

func (opt *baseOptions[T]) GetAdjustFlag() adjustFlag { return opt.adjustFlag }

// Currency 设置币种
func (opt *baseOptions[T]) Currency(curr currency) *T {
	currOpt := fmt.Sprintf("CurType=%d", curr)

	if optIdx := opt.findOptIdx("CurType"); optIdx < 0 {
		opt.options = append(opt.options, currOpt)
	} else {
		opt.options[optIdx] = currOpt
	}

	opt.currType = curr
	return opt.specificOpt
}

func (opt *baseOptions[T]) GetCurrency() currency { return opt.currType }

// BondType 设置债券价格模式
func (opt *baseOptions[T]) BondType(bond bondType) *T {
	bondOpt := fmt.Sprintf("Type=%d", bond)

	if optIdx := opt.findOptIdx("Type"); optIdx < 0 {
		opt.options = append(opt.options, bondOpt)
	} else {
		opt.options[optIdx] = bondOpt
	}

	opt.bondType = bond
	return opt.specificOpt
}

func (opt *baseOptions[T]) GetBondType() bondType { return opt.bondType }

// DateASC 设置日期升序
func (opt *baseOptions[T]) DateASC() *T {
	if optIdx := opt.findOptIdx("Order"); optIdx < 0 {
		opt.options = append(opt.options, "Order=1")
	} else {
		opt.options[optIdx] = "Order=1"
	}

	opt.dateDESC = false
	return opt.specificOpt
}

// DateDESC 设置日期降序
func (opt *baseOptions[T]) DateDESC() *T {
	if optIdx := opt.findOptIdx("Order"); optIdx < 0 {
		opt.options = append(opt.options, "Order=2")
	} else {
		opt.options[optIdx] = "Order=2"
	}

	opt.dateDESC = true
	return opt.specificOpt
}

func (opt *baseOptions[T]) IsDESC() bool { return opt.dateDESC }

// FillData 设置数据填充
func (opt *baseOptions[T]) FillData() *T {
	if optIdx := opt.findOptIdx("filldata"); optIdx < 0 {
		opt.options = append(opt.options, "filldata=1")
	} else {
		opt.options[optIdx] = "filldata=1"
	}

	opt.fillFlag = true

	return opt.specificOpt
}

func (opt *baseOptions[T]) IsFill() bool { return opt.fillFlag }

type DateOption[T Option] interface {
	BaseOption[T]

	ReportDate(time.Time) *T
	TradeDate(time.Time) *T
	StartDate(time.Time) *T
	EndDate(time.Time) *T
}

type dateOptions[T Option] struct {
	baseOptions[T]

	reportDate time.Time
	tradeDate  time.Time
	startDate  time.Time
	endDate    time.Time
}

func (opt *dateOptions[T]) format(wr io.Writer) {
	opt.baseOptions.format(wr)

	if opt.reportDate.IsZero() {
		fmt.Fprint(wr, " ReportDate:'not specified' ")
	} else {
		fmt.Fprintf(wr, " ReportDate:%+v ", opt.reportDate)
	}

	if opt.tradeDate.IsZero() {
		fmt.Fprint(wr, "TradeDate:'not specified' ")
	} else {
		fmt.Fprintf(wr, "TradeDate:%+v ", opt.tradeDate)
	}

	if opt.tradeDate.IsZero() {
		fmt.Fprint(wr, "StartDate:'not specified' ")
	} else {
		fmt.Fprintf(wr, "StartDate:%+v ", opt.startDate)
	}

	if opt.tradeDate.IsZero() {
		fmt.Fprint(wr, "EndDate:'not specified'")
	} else {
		fmt.Fprintf(wr, "EndDate:%+v", opt.endDate)
	}
}

// ReportDate 将指定日期向前对齐到季月最后一天
//
// 如已为季月最后一天，不进行处理
func (opt *dateOptions[T]) ReportDate(v time.Time) *T {
	rptDate := &v

	switch v.Month() {
	case time.March:
		if v.Day() == 31 {
			break
		}

		fallthrough
	case time.January, time.February:
		aligned := time.Date(
			v.Year()-1, time.December, 31, 0, 0, 0, 0, time.Local,
		)
		slog.Info(
			"report date align to season end",
			slog.Time("origin", v), slog.Time("align", aligned),
		)
		rptDate = &aligned
	case time.June:
		if v.Day() == 30 {
			break
		}

		fallthrough
	case time.April, time.May:
		aligned := time.Date(
			v.Year(), time.March, 31, 0, 0, 0, 0, time.Local,
		)
		slog.Info(
			"report date align to season end",
			slog.Time("origin", v), slog.Time("align", aligned),
		)
		rptDate = &aligned
	case time.September:
		if v.Day() == 30 {
			break
		}

		fallthrough
	case time.July, time.August:
		aligned := time.Date(
			v.Year(), time.June, 30, 0, 0, 0, 0, time.Local,
		)
		slog.Info(
			"report date align to season end",
			slog.Time("origin", v), slog.Time("align", aligned),
		)
		rptDate = &aligned
	case time.December:
		if v.Day() == 31 {
			break
		}

		fallthrough
	case time.October, time.November:
		aligned := time.Date(
			v.Year(), time.September, 30, 0, 0, 0, 0, time.Local,
		)
		slog.Info(
			"report date align to season end",
			slog.Time("origin", v), slog.Time("align", aligned),
		)
		rptDate = &aligned
	}

	optStr := fmt.Sprintf("ReportDate=%s", rptDate.Format("2006-01-02"))

	if optIdx := opt.findOptIdx("ReportDate"); optIdx < 0 {
		opt.options = append(opt.options, optStr)
	} else {
		opt.options[optIdx] = optStr
	}

	opt.reportDate = *rptDate
	return opt.specificOpt
}

func (opt *dateOptions[T]) GetReportDate() time.Time { return opt.reportDate }

// TradeDate 设置交易日期
func (opt *dateOptions[T]) TradeDate(v time.Time) *T {
	optStr := fmt.Sprintf("TradeDate=%s", v.Format("2006-01-02"))

	if optIdx := opt.findOptIdx("TradeDate"); optIdx < 0 {
		opt.options = append(opt.options, optStr)
	} else {
		opt.options[optIdx] = optStr
	}

	opt.tradeDate = v
	return opt.specificOpt
}

func (opt *dateOptions[T]) GetTradeDate() time.Time { return opt.tradeDate }

// StartDate 设置起始日期
func (opt *dateOptions[T]) StartDate(start time.Time) *T {
	startStr := fmt.Sprintf("StartDate=%s", start.Format("2006-01-02"))

	if optIdx := opt.findOptIdx("StartDate"); optIdx < 0 {
		opt.options = append(opt.options, startStr)
	} else {
		opt.options[optIdx] = startStr
	}

	opt.startDate = start
	return opt.specificOpt
}

func (opt *dateOptions[T]) GetStartDate() time.Time { return opt.startDate }

// EndDate 设置截止日期
func (opt *dateOptions[T]) EndDate(end time.Time) *T {
	endStr := fmt.Sprintf("EndDate=%s", end.Format("2006-01-02"))

	if optIdx := opt.findOptIdx("EndDate"); optIdx < 0 {
		opt.options = append(opt.options, endStr)
	} else {
		opt.options[optIdx] = endStr
	}

	opt.endDate = end
	return opt.specificOpt
}

func (opt *dateOptions[T]) GetEndDate() time.Time { return opt.endDate }

type KwOption[T Option] interface {
	DateOption[T]

	SetOption(string, any) (*T, error)
}

type kwOptions[T Option] struct {
	dateOptions[T]

	kwOptions map[string]any
}

func (opt *kwOptions[T]) format(buff io.Writer, skip ...string) {
	opt.dateOptions.format(buff)

	if len(opt.kwOptions) <= 0 {
		return
	}

	for k, v := range opt.kwOptions {
		if slices.Contains(skip, k) {
			continue
		}
		fmt.Fprintf(buff, " %s:%+v", k, v)
	}
}

func (opt *kwOptions[T]) SetOption(key string, value any) (*T, error) {
	if opt.kwOptions == nil {
		opt.kwOptions = make(map[string]any)
	}

	optStr := fmt.Sprintf("%s=%v", key, value)

	if optIdx := opt.findOptIdx(key); optIdx < 0 {
		opt.options = append(opt.options, optStr)
	} else {
		opt.options[optIdx] = optStr
	}

	opt.kwOptions[key] = value

	return opt.specificOpt, nil
}
