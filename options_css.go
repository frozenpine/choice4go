package choice4go

import (
	"github.com/valyala/bytebufferpool"
)

type CSSOptions struct {
	kwOptions[CSSOptions]
}

func (opt CSSOptions) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("CssOptions{")
	opt.format(buff)
	buff.WriteString("}")

	return buff.String()
}

func NewCssOptions() *CSSOptions {
	opt := new(CSSOptions)

	return opt.initBase(opt)
}

func init() {
	var _ KwOption[CSSOptions] = &CSSOptions{}
}
