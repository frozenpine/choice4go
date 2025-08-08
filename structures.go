package choice4go

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/valyala/bytebufferpool"
)

type Value interface {
	comparable
	~uint8 | ~uint | ~int | ~uint32 | ~int32 |
		~uint64 | ~int64 | ~float32 | ~float64 | ~[]uint8
}

var (
	dataPool  = sync.Pool{New: func() any { return &EQData{} }}
	valuePool = sync.Pool{New: func() any { return &EQValue{} }}
)

type EQValue struct {
	valueType   eqValueType
	valueBuffer [8]uint8
	valueString string
}

func (v *EQValue) Valid() bool {
	return v.valueType != ValueNull
}

func (v *EQValue) GetType() eqValueType {
	return v.valueType
}

func (v *EQValue) GetChar() uint8 {
	return v.valueBuffer[0]
}

func (v *EQValue) GetByte() byte {
	return v.valueBuffer[0]
}

func (v *EQValue) GetBool() bool {
	return binary.LittleEndian.Uint32(v.valueBuffer[:]) > 0
}

func (v *EQValue) GetShort() int16 {
	return int16(binary.LittleEndian.Uint16(v.valueBuffer[:]))
}

func (v *EQValue) GetUShort() uint16 {
	return binary.LittleEndian.Uint16(v.valueBuffer[:])
}

func (v *EQValue) GetInt() int {
	return int(binary.LittleEndian.Uint32(v.valueBuffer[:]))
}

func (v *EQValue) GetUInt() uint {
	return uint(binary.LittleEndian.Uint32(v.valueBuffer[:]))
}

func (v *EQValue) GetInt64() int64 {
	return int64(binary.LittleEndian.Uint64(v.valueBuffer[:]))
}

func (v *EQValue) GetUInt64() uint64 {
	return binary.LittleEndian.Uint64(v.valueBuffer[:])
}

func (v *EQValue) GetSingle() float32 {
	return math.Float32frombits(
		binary.LittleEndian.Uint32(v.valueBuffer[:]),
	)
}

func (v *EQValue) GetDouble() float64 {
	return math.Float64frombits(
		binary.LittleEndian.Uint64(v.valueBuffer[:]),
	)
}

func (v *EQValue) GetBytes() []byte {
	return v.valueBuffer[:]
}

func (v *EQValue) GetString() string {
	return v.valueString
}

func (v *EQValue) GetValue() any {
	switch v.valueType {
	case ValueNull:
		return nil
	case ValueChar:
		return v.GetChar()
	case ValueBool:
		return v.GetBool()
	case ValueShort:
		return v.GetShort()
	case ValueUShort:
		return v.GetUShort()
	case ValueInt:
		return v.GetInt()
	case ValueUInt:
		return v.GetUInt()
	case ValueInt64:
		return v.GetInt64()
	case ValueUInt64:
		return v.GetUInt64()
	case ValueSingle:
		return v.GetSingle()
	case ValueDouble:
		return v.GetDouble()
	case ValueBytes:
		return v.GetBytes()
	case ValueString:
		return v.GetString()
	default:
		slog.Error(
			"unknown value type",
			slog.Any("type", v.valueType),
		)

		return nil
	}
}

type Indicator struct {
	Code      string
	Date      time.Time
	sortedKey []string
	value     map[string]*EQValue
}

func (v Indicator) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("{Code:")
	buff.WriteString(v.Code)
	buff.WriteString(" Date:")
	buff.WriteString(v.Date.Format("2006-01-02"))
	buff.WriteString(" Indicators:{")
	for idx, name := range v.sortedKey {
		if idx > 0 {
			buff.WriteByte(' ')
		}
		buff.WriteString(
			fmt.Sprintf("%s:%+v", name, v.value[name].GetValue()),
		)
	}
	buff.WriteString("}}")

	return buff.String()
}

type EQData struct {
	codes      []string
	indicators []string
	dateList   []string
	values     []*EQValue
}

func (data *EQData) Iter() func(yield func(int, Indicator) bool) {
	codeSize := len(data.codes)
	indicatorSize := len(data.indicators)

	return func(yield func(int, Indicator) bool) {
		rowIdx := 0

		for idxDate, dateStr := range data.dateList {
			dateV := strings.SplitN(dateStr, "/", 3)
			yearV, err := strconv.Atoi(dateV[0])
			if err != nil {
				slog.Error(
					"parse date failed",
					slog.Any("error", err),
					slog.String("date", dateStr),
				)
			}
			monthV, err := strconv.Atoi(dateV[1])
			if err != nil {
				slog.Error(
					"parse date failed",
					slog.Any("error", err),
					slog.String("date", dateStr),
				)
			}
			dayV, err := strconv.Atoi(dateV[2])
			if err != nil {
				slog.Error(
					"parse date failed",
					slog.Any("error", err),
					slog.String("date", dateStr),
				)
			}

			date := time.Date(
				yearV, time.Month(monthV), dayV,
				0, 0, 0, 0, time.Local,
			)

			if err != nil {
				slog.Error(
					"parse date failed",
					slog.Any("error", err),
				)

				return
			}

			for idxCode, code := range data.codes {
				value := Indicator{
					Code:      code,
					Date:      date,
					sortedKey: data.indicators,
					value:     make(map[string]*EQValue),
				}

				for idxIndicator, indicator := range data.indicators {
					idx := codeSize*indicatorSize*idxDate + indicatorSize*idxCode + idxIndicator
					value.value[indicator] = data.values[idx]
				}

				if !yield(rowIdx, value) {
					return
				}

				rowIdx++
			}
		}
	}
}

type EQMsg struct {
	Version   int
	MsgType   eqMsgType
	RequestID int
	SerialID  int
	Data      *EQData
}
