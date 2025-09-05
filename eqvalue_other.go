//go:build !windows

package choice4go

func (v *EQValue) GetString() string {
	return v.valueString
}
