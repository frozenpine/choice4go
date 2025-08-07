package choice4go

import "encoding/binary"

type EQValue struct {
	ValueType   EQValueType
	valueBuffer [8]uint8
	// eqChar      string
}

func (v *EQValue) Valid() bool {
	return v.ValueType != ValueNull
}

func (v *EQValue) GetChar() uint8 {
	return v.valueBuffer[0]
}

func (v *EQValue) GetByte() uint8 {
	return v.valueBuffer[0]
}

func (v *EQValue) GetBool() bool {
	return binary.LittleEndian.Uint32(v.valueBuffer[:]) > 0
}

func (v *EQValue) GetShort() int16 {
	return int16(binary.LittleEndian.Uint16(v.valueBuffer[:]))
}

type EQData struct {
	Codes      []string
	Indicators []string
	Date       []string
	Values     []EQValue
}

type EQMsg struct {
	Version   int
	MsgType   EQMsgType
	RequestID int
	SerialID  int
	Data      *EQData
}
