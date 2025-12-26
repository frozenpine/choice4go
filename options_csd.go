package choice4go

import (
	"github.com/valyala/bytebufferpool"
)

type CSDOptions struct {
	kwOptions[CSDOptions]
}

func (opt CSDOptions) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("CsdOptions{")
	opt.format(buff)
	buff.WriteString("}")

	return buff.String()
}

func NewCsdOptions() *CSDOptions {
	opt := new(CSDOptions)

	return opt.initBase(opt)
}
