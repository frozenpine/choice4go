package choice4go

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/valyala/bytebufferpool"
)

//go:generate stringer -type isp -linecomment
type isp uint8

const (
	ISP_Auto isp = 0 // 未选择
	ISP_CT   isp = 1 // 电信
	ISP_CM   isp = 2 // 移动
	ISP_CU   isp = 3 // 联通
)

type startOptions struct {
	baseOptions

	testLatency bool
	forceLogin  bool
	recordLogin bool
	logLevel    slog.Level
	smsPhone    string
	timeout     time.Duration
	isp         isp
	useInnerNet bool
}

func (opt *startOptions) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("StartOptions{")
	fmt.Fprintf(buff, "TestLatency:%+v ", opt.testLatency)
	fmt.Fprintf(buff, "ForceLogin:%+v ", opt.forceLogin)
	fmt.Fprintf(buff, "RecordLogin:%+v ", opt.recordLogin)
	fmt.Fprintf(buff, "LogLevel:%+v ", opt.logLevel)
	fmt.Fprintf(buff, "SmsLogin:%+v ", opt.smsPhone)
	fmt.Fprintf(buff, "Timeout:%+v ", opt.timeout)
	fmt.Fprintf(buff, "ISP:%+v ", opt.isp)
	fmt.Fprintf(buff, "UseInnerNet:%+v}", opt.useInnerNet)

	return buff.String()
}

func NewStartOptions() *startOptions {
	return &startOptions{
		timeout: time.Second * 15,
	}
}

func (opt *startOptions) TestLatency() *startOptions {
	if optIdx := opt.findOptIdx("TestLatency"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, "TestLatency=1")
	} else {
		opt.baseOptions[optIdx] = "TestLatency=1"
	}

	opt.testLatency = true
	return opt
}

func (opt *startOptions) ForceLogin() *startOptions {
	if optIdx := opt.findOptIdx("ForceLogin"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, "ForceLogin=1")
	} else {
		opt.baseOptions[optIdx] = "ForceLogin=1"
	}

	opt.forceLogin = true
	return opt
}

func (opt *startOptions) RecordLoginInfo() *startOptions {
	if optIdx := opt.findOptIdx("RecordLoginInfo"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, "RecordLoginInfo=1")
	} else {
		opt.baseOptions[optIdx] = "RecordLoginInfo=1"
	}
	opt.recordLogin = true
	return opt
}

func (opt *startOptions) LogLevel(lvl slog.Level) *startOptions {
	var optStr string

	if lvl <= slog.LevelDebug {
		optStr = "LogLevel=1"
	} else if lvl <= slog.LevelInfo {
		optStr = "LogLevel=2"
	} else {
		optStr = "LogLevel=3"
	}

	if optIdx := opt.findOptIdx("LogLevel"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, optStr)
	} else {
		opt.baseOptions[optIdx] = optStr
	}

	opt.logLevel = lvl

	return opt
}

func (opt *startOptions) LoginSMS(phone string) *startOptions {
	if optIdx := opt.findOptIdx("LoginMode"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, "LoginMode=SXDL")
	} else {
		opt.baseOptions[optIdx] = "LoginMode=SXDL"
	}

	phoneOpt := fmt.Sprintf("PhoneNumber=%s", phone)

	if optIdx := opt.findOptIdx("PhoneNumber"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, phoneOpt)
	} else {
		opt.baseOptions[optIdx] = phoneOpt
	}

	opt.smsPhone = phone

	return opt
}

func (opt *startOptions) Timeout(d time.Duration) *startOptions {
	timeOpt := fmt.Sprintf("HTTPTimeout=%.0f", d.Seconds())

	if optIdx := opt.findOptIdx("HTTPTimeout"); optIdx < 0 {
		opt.baseOptions = append(opt.baseOptions, timeOpt)
	} else {
		opt.baseOptions[optIdx] = timeOpt
	}

	opt.timeout = d

	return opt
}

func (opt *startOptions) SelectISP(vender isp) *startOptions {
	if opt.findOptIdx("UseProxy", "UseInnerNet") >= 0 {
		slog.Warn(
			"isp selector conflict with UseProxy or UseInnerNet",
		)
	} else {
		switch vender {
		case ISP_CM, ISP_CT, ISP_CU:
			ispOpt := fmt.Sprintf("USEHTTP=%d", vender)

			if optIdx := opt.findOptIdx("USEHTTP"); optIdx < 0 {
				opt.baseOptions = append(opt.baseOptions, ispOpt)
			} else {
				opt.baseOptions[optIdx] = ispOpt
			}

			opt.isp = vender
		default:
			slog.Warn(
				"unknown ISP vender",
				slog.String("vender", vender.String()),
			)
		}
	}

	return opt
}

func (opt *startOptions) UseInnerNet() *startOptions {
	if opt.findOptIdx("USEHTTP") >= 0 {
		slog.Warn(
			"inner net option conflict with SelectISP",
		)
	} else {
		if optIdx := opt.findOptIdx("UseInnerNet"); optIdx < 0 {
			opt.baseOptions = append(opt.baseOptions, "UseInnerNet=1")
		} else {
			opt.baseOptions[optIdx] = "UseInnerNet=1"
		}

		opt.useInnerNet = true
	}

	return opt
}
