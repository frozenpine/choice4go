package choice4go

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:generate stringer -type eqMsgType -linecomment
type eqMsgType uint8

const (
	MsgTypeErr        eqMsgType = iota // 错误消息
	MsgTypeRsp                         // 应答
	MsgTypePartialRsp                  // 部分应答
	MsgTypeOther                       // 其他
)

//go:generate stringer -type cfnMode -linecomment
type cfnMode uint8

const (
	CfnStartToEnd cfnMode = 1 // 开始到结束中间的所有资讯
	CfgEndCound   cfnMode = 2 // 到结束时间的近count条资讯
)

//go:generate stringer -type eqValueType -linecomment
type eqValueType uint8

const (
	ValueNull   eqValueType = iota // Null
	ValueChar                      // Char
	ValueBool                      // Bool
	ValueShort                     // Short
	ValueUShort                    // Ushort
	ValueInt                       // Int
	ValueUInt                      // Uint
	ValueInt64                     // Int64
	ValueUInt64                    // Uint64
	ValueSingle                    // Single
	ValueDouble                    // Double
	ValueBytes                     // Bytes
	ValueString                    // String
)

const ValueByte = ValueChar

type StrArg string

func (arg StrArg) String() string {
	return string(arg)
}

func (arg StrArg) IsEmpty() bool {
	return arg == ""
}

type SliceArg []string

func (args SliceArg) String() string {
	return strings.Join(args, ",")
}

func (args SliceArg) IsEmpty() bool {
	return len(args) <= 0
}

type DateArg struct {
	time.Time
}

func (arg DateArg) String() string {
	return arg.Format("2006-01-02")
}

func (arg DateArg) IsEmpty() bool {
	return arg.IsZero()
}

func NewDateArg(year, month, day int) DateArg {
	return DateArg{time.Date(
		year, time.Month(month), day,
		0, 0, 0, 0, time.Local,
	)}
}

type argType interface {
	StrArg | SliceArg | DateArg

	fmt.Stringer
	IsEmpty() bool
}

//go:generate stringer -type funcType -linecomment
type funcType uint8

func (t funcType) OptionString() string {
	return fmt.Sprintf("FunType=%s", t.String())
}

func (t funcType) GetOption(key string) string {
	if key != "FunType" {
		return ""
	}

	return t.String()
}

func (t funcType) GetOptions(keys ...string) []string {
	if slices.Contains(keys, "FunType") {
		return []string{t.GetOption("FunType")}
	} else {
		return nil
	}
}

const (
	FuncCSD  funcType = iota // CSD
	FuncCSS                  // CSS
	FuncCSES                 // CSEC
)

//go:generate stringer -type ReportName -linecomment
type ReportName uint8

func (n ReportName) Name() string {
	switch n {
	case RptNameBalance:
		return "BalanceStatementSHSZ"
	case RptNameIncome:
		return "IncomeStatementSHSZ"
	case RptNameCashFlow:
		return "CashFlowStatementSHSZ"
	case RptNamePrediction:
		return "InstPredictionInfo"
	default:
		return n.String()
	}
}

func (n *ReportName) UnmarshalText(txt []byte) error {
	switch strings.ToLower(string(txt)) {
	case "balance", "BalanceStatementSHSZ":
		*n = RptNameBalance
	case "income", "IncomeStatementSHSZ":
		*n = RptNameIncome
	case "cashflow", "cash_flow", "CashFlowStatementSHSZ":
		*n = RptNameCashFlow
	case "predict", "prediction", "InstPredictionInfo":
		*n = RptNamePrediction
	default:
		return fmt.Errorf(
			"%w: unsupported value %s", ErrInvalidReportName, txt,
		)
	}

	return nil
}

const (
	RptNameInvalid    ReportName = iota // 未知报表名
	RptNameBalance                      // 资产负债表
	RptNameIncome                       // 利润表
	RptNameCashFlow                     // 现金流表
	RptNamePrediction                   // 盈利预测
)

//go:generate stringer -type ReportType -linecomment
type ReportType uint8

func (rptTyp *ReportType) UnmarshalText(txt []byte) error {
	switch strings.ToLower(string(txt)) {
	case "合并报表":
		*rptTyp = RptTypeCombined
	case "合并报表调整":
		*rptTyp = RptTypeCombAdjust
	case "母公司报表":
		*rptTyp = RptTypeParent
	case "母公司报表调整":
		*rptTyp = RptTypeParentAdjust
	default:
		v, err := strconv.Atoi(string(txt))

		if err != nil {
			return fmt.Errorf(
				"%w: unsupported value %s", ErrInvalidReportType, txt,
			)
		}

		*rptTyp = ReportType(v)
	}

	return nil
}

func (rptTyp *ReportType) Value() int {
	return int(*rptTyp)
}

const (
	RptTypeInvalid      ReportType = iota // 未知报表类型
	RptTypeCombined                       // 合并报表
	RptTypeCombAdjust                     // 合并报表调整
	RptTypeParent                         // 母公司报表
	RptTypeParentAdjust                   // 母公司报表调整
)

//go:generate stringer -type CompanyType -linecomment
type CompanyType uint8

const (
	CompanyInvalid   CompanyType = iota // 未知公司类型
	CompanyGeneral                      // 一般企业
	CompanyInsurance                    // 保险公司
	CompanyBank                         // 商业银行
	CompanySecurity                     // 证券公司
)

func (company *CompanyType) UnmarshalText(txt []byte) error {
	switch strings.ToLower(string(txt)) {
	case "general", "一般", "通用":
		*company = CompanyGeneral
	case "insurance", "保险":
		*company = CompanyInsurance
	case "bank", "银行":
		*company = CompanyBank
	case "security", "futures", "证券", "期货":
		*company = CompanySecurity
	default:
		v, err := strconv.Atoi(string(txt))
		if err != nil {
			return fmt.Errorf(
				"%w: unsupported value %s", ErrInvalidCompanyType, txt,
			)
		}

		*company = CompanyType(v)
	}

	return nil
}

func (v CompanyType) Name() string {
	switch v {
	case CompanyGeneral:
		return "通用"
	case CompanyInsurance:
		return "保险"
	case CompanyBank:
		return "银行"
	case CompanySecurity:
		return "证券"
	default:
		return v.String()
	}
}

//go:generate stringer -type RptSourceType -linecomment
type RptSourceType uint8

func (src *RptSourceType) UnmarshalText(txt []byte) error {
	switch strings.ToLower(string(txt)) {
	case "season1", "一季度报告", "一季报":
		*src = RptSrcSeason1
	case "season2", "半年度报告", "二季报", "半年报":
		*src = RptSrcSeason2
	case "season3", "三季度报告", "三季报":
		*src = RptSrcSeason3
	case "season4", "年报报告", "四季报", "年报":
		*src = RptSrcSeason4
	default:
		v, err := strconv.Atoi(string(txt))
		if err != nil {
			return fmt.Errorf(
				"%w: unsupported value %s", ErrInvalidSourceType, txt,
			)
		}

		*src = RptSourceType(v)
	}

	return nil
}

func (src RptSourceType) Name() string {
	switch src {
	case RptSrcSeason1:
		return "一季度报告"
	case RptSrcSeason2:
		return "半年度报告"
	case RptSrcSeason3:
		return "三季度报告"
	case RptSrcSeason4:
		return "年报报告"
	default:
		return src.String()
	}
}

const (
	RptSrcInvalid RptSourceType = iota // 未知报告来源
	RptSrcSeason1                      // 一季报
	RptSrcSeason2                      // 半年报
	RptSrcSeason3                      // 三季报
	RptSrcSeason4                      // 年报
)

type Boolean bool

func (b *Boolean) UnmarshalText(txt []byte) error {
	switch str := strings.ToLower(string(txt)); str {
	case "是", "真", "t", "true", "y", "yes":
		*b = true
	case "否", "假", "f", "false", "n", "no":
		*b = false
	default:
		v, err := strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf(
				"%w: unknown choice boolean value %s",
				ErrInvalidDataValue, str,
			)
		}

		if v == 0 {
			*b = false
		} else {
			*b = true
		}
	}

	return nil
}

func (b Boolean) Name() string {
	if b {
		return "是"
	} else {
		return "否"
	}
}

func (b Boolean) Value() bool {
	return bool(b)
}
