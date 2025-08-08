package choice4go

import (
	"fmt"

	"github.com/valyala/bytebufferpool"
)

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

type csdOptions struct {
	baseOptions

	period     period
	adjustFlag adjustFlag
	currType   currency
	bondType   bondType
	dateDESC   bool
}

func NewCsdOptions() *csdOptions {
	return &csdOptions{
		period:     Daily,
		adjustFlag: NoAdjusted,
		currType:   CurrOrigin,
		bondType:   BondClean,
	}
}

func (opt *csdOptions) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("CsdOptions{")
	fmt.Fprintf(buff, "Period:%+v ", opt.period)
	fmt.Fprintf(buff, "Adjust:%+v ", opt.adjustFlag)
	fmt.Fprintf(buff, "Currency:%+v ", opt.currType)
	fmt.Fprintf(buff, "BondPrice:%+v", opt.bondType)
	if opt.dateDESC {
		fmt.Fprintf(buff, "DateSort:DESC}")
	} else {
		fmt.Fprintf(buff, "DateSort:ASC}")
	}

	return buff.String()
}

func (opt *csdOptions) Period(p period) *csdOptions {
	periodOpt := fmt.Sprintf("Period=%d", p)

	if optIdx := opt.findOptIdx("Period"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, periodOpt)
	} else {
		opt.baseOptions[optIdx] = periodOpt
	}

	opt.period = p

	return opt
}

func (opt *csdOptions) Adjust(flag adjustFlag) *csdOptions {
	flagOpt := fmt.Sprintf("AdjustFlag=%d", flag)

	if optIdx := opt.findOptIdx("AdjustFlag"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, flagOpt)
	} else {
		opt.baseOptions[optIdx] = flagOpt
	}

	opt.adjustFlag = flag
	return opt
}

func (opt *csdOptions) Currency(curr currency) *csdOptions {
	currOpt := fmt.Sprintf("CurType=%d", curr)

	if optIdx := opt.findOptIdx("CurType"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, currOpt)
	} else {
		opt.baseOptions[optIdx] = currOpt
	}

	opt.currType = curr
	return opt
}

func (opt *csdOptions) BondType(bond bondType) *csdOptions {
	bondOpt := fmt.Sprintf("Type=%d", bond)

	if optIdx := opt.findOptIdx("Type"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, bondOpt)
	} else {
		opt.baseOptions[optIdx] = bondOpt
	}

	opt.bondType = bond
	return opt
}

func (opt *csdOptions) DateASC() *csdOptions {
	if optIdx := opt.findOptIdx("Order"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, "Order=1")
	} else {
		opt.baseOptions[optIdx] = "Order=1"
	}

	opt.dateDESC = false
	return opt
}

func (opt *csdOptions) DateDESC() *csdOptions {
	if optIdx := opt.findOptIdx("Order"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, "Order=2")
	} else {
		opt.baseOptions[optIdx] = "Order=2"
	}

	opt.dateDESC = true
	return opt
}
