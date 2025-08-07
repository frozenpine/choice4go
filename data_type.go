package choice4go

type EQMsgType uint8

const (
	MsgTypeErr EQMsgType = iota
	MsgTypeRsp
	MsgTypePartialRsp
	MsgTypeOther
)

type EQValueType uint8

const (
	ValueNull EQValueType = iota
	ValueChar
	ValueBool
	ValueShort
	ValueUShort
	ValueInt
	ValueUInt
	ValueInt64
	ValueUInt64
	ValueSingle
	ValueDouble
	ValueBytes
	ValueString
)

const ValueByte = ValueChar
