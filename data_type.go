package choice4go

//go:generate stringer -type eqMsgType -linecomment
type eqMsgType uint8

const (
	MsgTypeErr        eqMsgType = iota // 错误消息
	MsgTypeRsp                         // 应答
	MsgTypePartialRsp                  // 部分应答
	MsgTypeOther                       // 其他
)

//go:generate stringer -type eqValueType -linecomment
type eqValueType uint8

const (
	ValueNull   eqValueType = iota // Null
	ValueChar                      // Char
	ValueBool                      // Bool
	ValueShort                     // Short
	ValueUShort                    // Ushort
	ValueInt                       // Int
	ValueUInt                      // Uint
	ValueInt64                     // Int64
	ValueUInt64                    // Uint64
	ValueSingle                    // Single
	ValueDouble                    // Double
	ValueBytes                     // Bytes
	ValueString                    // String
)

const ValueByte = ValueChar
