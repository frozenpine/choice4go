package choice4go

import (
	"fmt"
	"slices"
	"strings"
)

type Option interface {
	fmt.Stringer

	OptionString() string
}

type baseOptions []string

func (opt baseOptions) OptionString() string {
	return strings.Join(opt, ",")
}

func (opt baseOptions) findOptIdx(prefix ...string) int {
	return slices.IndexFunc(opt, func(v string) bool {
		for _, match := range prefix {
			if strings.HasPrefix(v, match) {
				return true
			}
		}

		return false
	})
}
