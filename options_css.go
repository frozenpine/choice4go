package choice4go

import (
	"github.com/valyala/bytebufferpool"
)

type CSSOptions struct {
	kwOptions[CSSOptions]

	ttm ttmType
}

func (opt CSSOptions) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("CssOptions{")
	opt.format(buff)
	buff.WriteString("}")

	return buff.String()
}

// TTMType 设置TTM基准日
func (opt *CSSOptions) TTMType(ttm ttmType) {
	opt.ttm = ttm
}

func (opt *CSSOptions) GetTTM() ttmType { return opt.ttm }

func NewCssOptions() *CSSOptions {
	opt := new(CSSOptions)

	return opt.initBase(opt)
}
