package choice4go

import (
	"context"
	"fmt"
	"sync/atomic"
)

var (
	singleton atomic.Pointer[Choice]

	errSingletonInstance = fmt.Errorf(
		"%w: singleton instance not created", ErrInitialized,
	)
)

func GetInstance() *Choice {
	return singleton.Load()
}

func Start(
	ctx context.Context,
	user, pass string,
	options Option,
) (err error) {
	if ins := singleton.Load(); ins != nil {
		return ins.Start(ctx, user, pass, options)
	} else {
		return errSingletonInstance
	}
}

func Restart() error {
	if ins := singleton.Load(); ins != nil {
		return ins.Restart()
	} else {
		return errSingletonInstance
	}
}

func Stop() error {
	if ins := singleton.Load(); ins != nil {
		return ins.Stop()
	} else {
		return errSingletonInstance
	}
}

func Csd(
	codes, indicators SliceArg, start, end DateArg, options Option,
) (*EQData, error) {
	if ins := singleton.Load(); ins != nil {
		return ins.Csd(codes, indicators, start, end, options)
	} else {
		return nil, errSingletonInstance
	}
}

func Css(
	codes, indicators SliceArg, options Option,
) (*EQData, error) {
	if ins := singleton.Load(); ins != nil {
		return ins.Css(codes, indicators, options)
	} else {
		return nil, errSingletonInstance
	}
}

func CSes(
	blockCodes, indicators SliceArg, options Option,
) (*EQData, error) {
	if ins := singleton.Load(); ins != nil {
		return ins.CSes(blockCodes, indicators, options)
	} else {
		return nil, errSingletonInstance
	}
}

func TradeDates(
	start, end DateArg, options Option,
) (*EQData, error) {
	if ins := singleton.Load(); ins != nil {
		return ins.TradeDates(start, end, options)
	} else {
		return nil, errSingletonInstance
	}
}

func Sector(
	code StrArg, tradeDate DateArg, options Option,
) (*EQData, error) {
	if ins := singleton.Load(); ins != nil {
		return ins.Sector(code, tradeDate, options)
	} else {
		return nil, errSingletonInstance
	}
}

func Ctr(
	name StrArg, indicators SliceArg, options Option,
) (*EQCtrData, error) {
	if ins := singleton.Load(); ins != nil {
		return ins.Ctr(name, indicators, options)
	} else {
		return nil, errSingletonInstance
	}
}

func Edb(
	ids SliceArg, options Option,
) (*EQData, error) {
	if ins := singleton.Load(); ins != nil {
		return ins.Edb(ids, options)
	} else {
		return nil, errSingletonInstance
	}
}

func EdbQuery(
	ids, indicators SliceArg, options Option,
) (*EQData, error) {
	if ins := singleton.Load(); ins != nil {
		return ins.EdbQuery(ids, indicators, options)
	} else {
		return nil, errSingletonInstance
	}
}

func Cfn(
	codes SliceArg, content StrArg, mode cfnMode, options Option,
) (*EQData, error) {
	if ins := singleton.Load(); ins != nil {
		return ins.Cfn(codes, content, mode, options)
	} else {
		return nil, errSingletonInstance
	}
}

func CfnQuery(
	options Option,
) (*EQData, error) {
	if ins := singleton.Load(); ins != nil {
		return ins.CfnQuery(options)
	} else {
		return nil, errSingletonInstance
	}
}
