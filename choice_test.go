package choice4go

import (
	"context"
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"
)

var (
	libDir  = "./libs"
	libName = "EMQuantAPI"
	cfgDir  = "./cfg"
	user    = ""
	pass    = ""
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func TestChoiceCsd(t *testing.T) {
	choice, err := NewChoice(
		libDir, libName, cfgDir,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = choice.Start(
		context.TODO(), user, pass,
		NewStartOptions().
			TestLatency().
			ForceLogin().
			LogLevel(slog.LevelDebug),
	); err != nil {
		t.Fatal(err)
	}
	defer choice.Stop()

	if results, err := choice.Csd(
		[]string{"AG0.SHF", "600519.SH"},
		[]string{
			"OPEN", "CLOSE", "HIGH", "LOW", "VOLUME", "AMOUNT", "PRECLOSE",
			"CHANGE", "MAINFORCE",
		},
		NewDateArg(2024, 1, 1), NewDateArg(2024, 12, 31),
		nil,
	); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range results.Iter() {
			t.Logf("%+v", v)
		}
	}

	// if results, err := choice.TradeDates(
	// 	NewDateArg(2024, 1, 1), NewDateArg(2024, 12, 31),
	// 	nil,
	// ); err != nil {
	// 	t.Fatal(err)
	// } else {
	// 	for _, v := range results.Iter() {
	// 		t.Logf("%+v", v)
	// 	}
	// }

	// if results, err := choice.Csd(
	// 	[]string{"000002.SZ", "300059.SZ"},
	// 	[]string{"OPEN", "HIGH", "LOW", "CLOSE"},
	// 	NewDateArg(2016, 1, 10), NewDateArg(2016, 4, 13),
	// 	NewCsdOptions().
	// 		Period(Daily).
	// 		Adjust(NoAdjusted).
	// 		Currency(CurrCNY).
	// 		BondType(BondDirty),
	// ); err != nil {
	// 	t.Fatal(err)
	// } else {
	// 	for _, v := range results.Iter() {
	// 		t.Logf("%+v", v)
	// 	}
	// }
}

func TestChoiceCsdPredict(t *testing.T) {
	choice, err := NewChoice(
		libDir, libName, cfgDir,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = choice.Start(
		context.TODO(), user, pass,
		NewStartOptions().
			TestLatency().
			ForceLogin().
			LogLevel(slog.LevelDebug),
	); err != nil {
		t.Fatal(err)
	}
	defer choice.Stop()

	symbols := []string{
		"002961.SZ", "600519.SH",
	}

	if results, err := choice.Csd(
		symbols,
		[]string{
			"RATINGAVG", "RATINGAVGCHN", "RATINGAVGENG", "RATINGINSTNUM",
			"RATINGMAINTAIN", "RATINGUPGRADE", "RATINGDOWNGRADE",
			"RATINGNUMOFBUY", "RATINGNUMOFOUTPERFORM", "RATINGNUMOFHOLD",
			"RATINGNUMOFUNDERPERFORM", "RATINGNUMOFSELL", "UPGRADE",
		},
		NewDateArg(2024, 1, 1), NewDateArg(2025, 8, 25),
		nil,
	); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range results.Iter() {
			t.Logf("%+v", v)
		}
	}
}

func TestChoiceCfc(t *testing.T) {
	choice, err := NewChoice(
		libDir, libName, cfgDir,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = choice.Start(
		context.TODO(), user, pass,
		NewStartOptions().
			// TestLatency().
			ForceLogin().
			LogLevel(slog.LevelDebug),
	); err != nil {
		t.Fatal(err)
	}
	defer choice.Stop()

	if results, err := choice.Cfc(
		[]string{"002961.SZ"},
		[]string{
			"EPSBASIC", "EPSDILUTED", "EPSDILUTEDEND", "EPSDILUTEDNEW",
			"EPSEXBASIC", "EPSEXDILUTED", "EPSEXDILUTEDEND", "EPSEXDILUTEDNEW",
			"EPSTTM", "EPSNEW", "BPS", "BPSDILUTEDNEW", "BPSNEW", "CFOPS",
			"CFOPSTTM", "CFOPSDILUTEDNEW", "GRPS", "ORPS", "ORPSTTM",
			"CAPITALRESERVEPS", "CAPITALRESERVEPSNEW", "SURPLUSRESERVEPS",
			"UNDISTRIBUTEDPS", "UNDISTRIBUTEDPSN", "RETAINEDPS", "CFPS",
			"CFPSTTM", "CFPSDILUTEDNEW", "EBITPS", "FCFFPS", "FCFEPS", "EBITDAPS",
		},
		FuncCSS,
	); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range results.Iter() {
			t.Logf("%+v", v)
		}
	}
}

func TestChoiceCss(t *testing.T) {
	choice, err := NewChoice(
		libDir, libName, cfgDir,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = choice.Start(
		context.TODO(), user, pass,
		NewStartOptions().
			// TestLatency().
			ForceLogin().
			LogLevel(slog.LevelDebug),
	); err != nil {
		t.Fatal(err)
	}
	defer choice.Stop()

	rptDate := time.Date(2025, 6, 30, 0, 0, 0, 0, time.Local)
	tradeDate := time.Date(2025, 8, 25, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2025, 8, 25, 0, 0, 0, 0, time.Local)

	symbols := []string{
		"002961.SZ", "600519.SH",
	}

	if results, err := choice.Css(
		symbols,
		[]string{
			"EPSBASIC", "EPSDILUTED", "EPSDILUTEDEND", "EPSDILUTEDNEW",
			"EPSEXBASIC", "EPSEXDILUTED", "EPSEXDILUTEDEND", "EPSEXDILUTEDNEW",
			"EPSTTM", "EPSNEW", "BPS", "BPSDILUTEDNEW", "BPSNEW", "CFOPS",
			"CFOPSTTM", "CFOPSDILUTEDNEW", "GRPS", "ORPS", "ORPSTTM",
			"CAPITALRESERVEPS", "CAPITALRESERVEPSNEW", "SURPLUSRESERVEPS",
			"UNDISTRIBUTEDPS", "UNDISTRIBUTEDPSN", "RETAINEDPS", "CFPS",
			"CFPSTTM", "CFPSDILUTEDNEW", "EBITPS", "FCFFPS", "FCFEPS", "EBITDAPS",
		},
		NewCssOptions().
			ReportDate(rptDate).
			EndDate(endDate).
			TradeDate(tradeDate),
	); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range results.Iter() {
			t.Logf("%+v", v)
		}
	}
}

func TestChoiceCtr(t *testing.T) {
	choice, err := NewChoice(
		libDir, libName, cfgDir,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = choice.Start(
		context.TODO(), user, pass,
		NewStartOptions().
			// TestLatency().
			ForceLogin().
			LogLevel(slog.LevelDebug),
	); err != nil {
		t.Fatal(err)
	}
	defer choice.Stop()

	// BalanceStatementSHSZ 资产负债表
	// IncomeStatementSHSZ 利润表
	// CashFlowStatementSHSZ 现金流表
	// InstPredictionInfo 盈利预测
	// Options:ReportType 1 合并报表
	// Options:ReportType 2 合并报表调整
	// Options:ReportType 3 母公司报表
	// Options:ReportType 3 母公司报表调整
	// if result, err := choice.Ctr(
	// 	StrArg(RptNameBalance.Name()),
	// 	[]string{
	// 		"MONETARYFUND", "SETTLEMENTPROVISION", "LENDFUND",
	// 		"TRADE_FINASSET_NOTFVTPL", "MARGINOUTFUND", "DERIVEFASSET",
	// 		"ACCOUNTBILLREC", "BILLREC", "ACCOUNTREC", "FINANCE_RECE",
	// 		"ADVANCEPAY", "PREMIUMREC", "RIREC", "RICONTACTRESERVEREC",
	// 		"TOTAL_OTHER_RECE", "INTERESTREC", "DIVIDENDREC", "OTHERREC",
	// 		"EXPORTREBATEREC", "SUBSIDYREC", "INTERNALREC", "BUYSELLBACKFASSET",
	// 		"AMORCOSTFASSET", "FVALUECOMPFASSET", "INVENTORY", "CONTRACTASSET",
	// 		"CLHELDSALEASS", "NONLASSETONEYEAR", "DLYWZC", "OTHERLASSET",
	// 		"LASSETOTHER", "LASSETBALANCE", "SUMLASSET", "LOANADVANCES",
	// 		"CREDINV", "AMORCOSTFASSETFLD", "OTHCREDINV", "FVALUECOMPFASSETFLD",
	// 		"SALEABLEFASSET", "HELDMATURITYINV", "LTREC", "LTEQUITYINV",
	// 		"OTHEREQUITYINV", "OTHERNONFASSET", "ESTATEINVEST", "FIXEDASSET",
	// 		"CONSTRUCTIONPROGRESS", "CONSTRUCTIONMATERIAL",
	// 		"LIQUIDATEFIXEDASSET", "PRODUCTBIOLOGYASSET", "OILGASASSET",
	// 		"USERIGHT_ASSET", "INTANGIBLEASSET", "DEVELOPEXP", "GOODWILL",
	// 		"LTDEFERASSET", "DEFERINCOMETAXASSET", "OTHERNONLASSET",
	// 		"NONLASSETOTHER", "NONLASSETBALANCE", "SUMNONLASSET",
	// 		"ASSETOTHER", "ASSETBALANCE", "SUMASSET",
	// 		"STR_COMBINETYPE",
	// 	},
	// 	NewCtrOptions().
	// 		ReportDate(time.Date(2025, 1, 31, 0, 0, 0, 0, time.Local)).
	// 		// Range(
	// 		// 	time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
	// 		// 	time.Date(2025, 12, 31, 0, 0, 0, 0, time.Local),
	// 		// ).
	// 		Code("601628.SH"),
	// ); err != nil {
	// 	t.Fatal(err)
	// } else {
	// 	for _, v := range result.Iter() {
	// 		t.Logf("%+v", v)
	// 	}
	// }

	if result, err := choice.Ctr(
		StrArg(RptNameCashFlow.Name()),
		[]string{
			"INTANDCOMMREC", "OTHEROPERATEREC", "OPERATEFLOWINOTHER",
			"OPERATEFLOWINBALANCE", "SUMOPERATEFLOWIN", "INTANDCOMMPAY",
			"EMPLOYEEPAY", "TAXPAY", "OTHEROPERATEPAY", "OPERATEFLOWOUTOTHER",
			"OPERATEFLOWOUTBALANCE", "SUMOPERATEFLOWOUT", "OPERATEFLOWOTHER",
			"OPERATEFLOWBALANCE", "NETOPERATECASHFLOW", "DISPOSALINVREC",
			"INVINCOMEREC", "DISPFILASSETREC", "DISPSUBSIDIARYREC",
			"OTHERINVREC", "INVFLOWINOTHER", "INVFLOWINBALANCE", "SUMINVFLOWIN",
			"BUYFILASSETPAY", "INVPAY", "OTHERINVPAY", "INVFLOWOUTOTHER",
			"INVFLOWOUTBALANCE", "SUMINVFLOWOUT", "INVFLOWOTHER",
			"NETINVCASHFLOW", "ACCEPTINVREC", "ISSUEBONDREC", "OTHERFINAREC",
			"FINAFLOWINOTHER", "FINAFLOWINBALANCE", "SUMFINAFLOWIN",
			"REPAYDEBTPAY", "DIVIPROFITORINTPAY", "OTHERFINAPAY",
			"FINAFLOWOUTOTHER", "FINAFLOWOUTBALANCE", "SUMFINAFLOWOUT",
			"FINAFLOWOTHER", "FINAFLOWBALANCE", "NETFINACASHFLOW",
			"EFFECTEXCHANGERATE", "NICASHEQUIOTHER", "NICASHEQUIBALANCE",
			"NICASHEQUI", "CASHEQUIBEGINNING", "CASHEQUIENDINGOTHER",
			"CASHEQUIENDINGBALANCE", "CASHEQUIENDING", "NETPROFIT",
			"ASSETDEVALUE", "FIXANDESTATEDEPR", "INTANGIBLEASSETAMOR",
			"LTDEFEREXPAMOR", "DISPFILASSETLOSS", "FVALUELOSS", "INVLOSS",
			"DEFERTAX", "DEFERTAXASSETREDUCE", "DEFERTAXLIABADD",
			"OPERATERECREDUCE", "OPERATEPAYADD", "OTHER", "DEC_JYHDCSDXJLLJEQT",
			"DEC_JYHDCSDXJLLJEPH", "DEC_JYHDCSDXJLLJE", "DEBTTOCAPITAL",
			"CBONEYEAR", "FINALEASEFIXEDASSET", "NOREFERCASHOTHER",
			"DEC_XJDQMYE", "DEC_XJDQCYE", "DEC_XJJZJECETS",
			"DEC_XJJZJECEHJ",
			"REPORTSOURCETYPE", "STR_COMBINETYPE",
			"OPINIONTYPE", "FIRSTNOTICEDATE",
		},
		NewCtrOptions().
			ReportDate(time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local)).
			// Range(
			// 	time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			// 	time.Date(2025, 12, 31, 0, 0, 0, 0, time.Local),
			// ).
			Code("600519.SH"),
	); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range result.Iter() {
			t.Logf("%+v", v)
		}
	}
}

func TestWinnerList(t *testing.T) {
	choice, err := NewChoice(
		libDir, libName, cfgDir,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = choice.Start(
		context.TODO(), user, pass,
		NewStartOptions().
			// TestLatency().
			ForceLogin().
			LogLevel(slog.LevelDebug),
	); err != nil {
		t.Fatal(err)
	}
	defer choice.Stop()

	codes := []string{
		"AG0.SHF",
	}

	// opt, err := NewCssOptions().
	// 	TradeDate(time.Date(2025, 9, 9, 0, 0, 0, 0, time.Local)).
	// 	SetOption("Rank", 0)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	opt := NewCssOptions().
		TradeDate(time.Date(2025, 9, 9, 0, 0, 0, 0, time.Local))

	// 	# 期货现货 持买单量进榜会员名称 持卖单量进榜会员名称 成交量进榜会员名称 成交量进榜会员合计 多单量进榜会员合计 空单量进榜会员合计
	// data=c.css("AG0.SHF","FTNEWLMEMNAME,FTNEWSMEMNAME,FTNEWVOLMEMNAME,FTNEWVOLMEMTOTAL,FTNEWLMEMTOTAL,FTNEWSMEMTOTAL","TradeDate=2025-09-10,Rank=0")
	// # 期货现货 持买单量 持买单量比上交易日增减 持卖单量 持卖单量比上交易日增减 成交量 成交量比上交易日增减
	// data=c.css("AG0.SHF","FTLONGNUM,FTLONGCHG,FTSHORTNUM,FTSHORTCHG,FTVOLUUME,FTVOLUMECHG","TradeDate=2025-09-10,Rank=1")

	for rank := range 20 {
		opt, err = opt.SetOption("Rank", rank+1)
		if err != nil {
			t.Fatal(err)
		}

		result, err := Css(
			codes,
			[]string{
				// csd
				// "FTLONGNUM", "FTSHORTNUM", "FTVOLUME", "FTLCHANGE", "FTSCHANGE",
				// "FTVCHANGE", "FTVCOUNT", "FTLCOUNT", "FTSCOUNT", "FTREGORDERVOL",
				"FTLONGNUM", "FTLONGCHG", "FTSHORTNUM", "FTSHORTCHG", "FTVOLUUME", "FTVOLUMECHG",
				// css
				"FTNEWLMEMNAME", "FTNEWSMEMNAME", "FTNEWVOLMEMNAME",
				"FTNEWVOLMEMTOTAL", "FTNEWLMEMTOTAL", "FTNEWSMEMTOTAL",
			},
			// NewDateArg(2025, 1, 1), NewDateArg(2025, 9, 8),
			opt,
		)

		if err != nil {
			t.Fatal(err)
		}

		for _, d := range result.Iter() {
			t.Log(d)
		}
	}
}

func TestContract(t *testing.T) {
	choice, err := NewChoice(
		libDir, libName, cfgDir,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = choice.Start(
		context.TODO(), user, pass,
		NewStartOptions().
			// TestLatency().
			ForceLogin().
			LogLevel(slog.LevelDebug),
	); err != nil {
		t.Fatal(err)
	}
	defer choice.Stop()

	codes := []string{
		"AG0.SHF", "CU0.SHF", "AL0.SHF", "ZN0.SHF", "SN0.SHF", "PB0.SHF",
		"NI0.SHF", "AO0.SHF", "AD0.SHF", "AU0.SHF", "RB0.SHF", "WR0.SHF",
		"HC0.SHF", "SS0.SHF", "FU0.SHF", "BU0.SHF", "BR0.SHF", "OP0.SHF",
		"RU0.SHF", "SP0.SHF",
	}

	// 	# 2025-09-10 14:27:04
	// # undefined 交易品种 交易单位 合约乘数 报价单位 最小变动价位 标准合约上市日 合约月份说明 交易时间说明 最后交易日说明 交割日期说明 最低交易保证金 标的代码
	// data=c.css("AG0.SHF","FTTRANSTYPE,FTTRANSUNIT,CONTRACTMUL,FTPRICEUNIT,FTMINPRICECHG,LISTDATE,FTCONTRADATEINTRO,FTTRANSDATEINSTRO,FTLTRANSDATE,FTDELIVDATEINTRO,FTFIRSTTRANSMARGIN,UNDERLYINGCODE","TRADEDATE=2025-09-10")

	indicators := []string{
		"FTTRANSTYPE", "FTTRANSUNIT", "CONTRACTMUL", "FTPRICEUNIT",
		"FTMINPRICECHG", "LISTDATE", "FTCONTRADATEINTRO", "FTTRANSDATEINSTRO",
		"FTLTRANSDATE", "FTDELIVDATEINTRO", "FTFIRSTTRANSMARGIN", "UNDERLYINGCODE",
	}

	opt := NewCssOptions().
		TradeDate(time.Date(2025, 9, 9, 0, 0, 0, 0, time.Local))

	result, err := Css(
		codes, indicators, opt,
	)

	if err != nil {
		t.Fatal(err)
	}

	for _, d := range result.Iter() {
		t.Log(d)
	}
}

func TestIndex(t *testing.T) {
	choice, err := NewChoice(
		libDir, libName, cfgDir,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = choice.Start(
		context.TODO(), user, pass,
		NewStartOptions().
			// TestLatency().
			ForceLogin().
			LogLevel(slog.LevelDebug),
	); err != nil {
		t.Fatal(err)
	}
	defer choice.Stop()

	codes := []string{
		// 中证1000
		"000852.SH",
		// 沪深300
		"000300.SH",
		// 中证500
		"000905.SH",
		// 上证50
		"000016.SH",
	}

	indicators := []string{
		"CODE", "SHORTNAME", "NAME", "BASICPOINT", "BASICDATE", "PUBLISHDATE",
		"DELISTDATE", "COMPONENTNUM", "FIRSTDAYOFCONSTITUENTS", "MAKERNAME",
		"INDEXPROFILE", "CALCULATION", "METHODOLOGY", "OFFICIALSTYLE",
		"SAMPLECHGPRIN", "INDEXTYPE", "TRACKEDBYFUNDS", "TARGETCODE",
	}

	opt := NewCssOptions().
		TradeDate(time.Date(2025, 9, 23, 0, 0, 0, 0, time.Local))

	result, err := Css(
		codes, indicators, opt,
	)

	if err != nil {
		t.Fatal(err)
	}

	for _, d := range result.Iter() {
		t.Log(d)
	}
}

func TestIndexConstituent(t *testing.T) {
	choice, err := NewChoice(
		libDir, libName, cfgDir,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err = choice.Start(
		context.TODO(), user, pass,
		NewStartOptions().
			// TestLatency().
			ForceLogin().
			LogLevel(slog.LevelDebug),
	); err != nil {
		t.Fatal(err)
	}
	defer choice.Stop()

	codes := []string{
		// 中证1000
		"000852.SH",
		// 沪深300
		"000300.SH",
		// 中证500
		"000905.SH",
		// 上证50
		"000016.SH",
	}

	indicators := []string{
		"INDEXCODE", "SECUCODE", "TRADEDATE", "NAME", "CLOSE", "PCTCHANGE",
		"WEIGHT", "CONTRIBUTEPT", "SHRMARKETVALUE", "MV", "TOTALTRADABLE",
		"SHARETOTAL",
	}
	opt := NewCtrOptions().EndDate(
		time.Date(2025, 9, 23, 0, 0, 0, 0, time.Local),
	)

	for _, idx := range codes {
		opt.SetOption("IndexCode", idx)

		if result, err := Ctr(
			"INDEXCONSTITUENT", indicators, opt,
		); err != nil {
			t.Fatal(err)
		} else {
			csvFile, err := os.Create(fmt.Sprintf("%s.csv", idx))
			if err != nil {
				t.Fatal(err)
			}
			wr := csv.NewWriter(csvFile)
			defer func() {
				wr.Flush()
				csvFile.Close()
			}()

			for idx, r := range result.Iter() {
				if idx == 0 {
					wr.Write(r.indicators)
				}

				values := make([]string, 0, len(r.indicators))
				for _, v := range r.Iter() {
					values = append(values, fmt.Sprintf("%+v", v.GetValue()))
				}
				wr.Write(values)

				t.Log(r)
			}
		}
	}
}
