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
	dataPool    = sync.Pool{New: func() any { return &EQData{} }}
	ctrDataPool = sync.Pool{New: func() any { return &EQCtrData{} }}
	valuePool   = sync.Pool{New: func() any { return &EQValue{} }}
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

func (v *EQValue) Char(c uint8) {
	v.valueType = ValueChar
	v.valueBuffer[0] = c
}

func (v *EQValue) GetByte() byte {
	return v.valueBuffer[0]
}

func (v *EQValue) Byte(c byte) {
	v.valueType = ValueByte
	v.valueBuffer[0] = c
}

func (v *EQValue) GetBool() bool {
	return binary.LittleEndian.Uint32(v.valueBuffer[:]) > 0
}

func (v *EQValue) Bool(b bool) {
	v.valueType = ValueBool
	if b {
		binary.LittleEndian.PutUint32(v.valueBuffer[:], 1)
	} else {
		binary.LittleEndian.PutUint32(v.valueBuffer[:], 0)
	}
}

func (v *EQValue) GetShort() int16 {
	return int16(binary.LittleEndian.Uint16(v.valueBuffer[:]))
}

func (v *EQValue) Short(s int16) {
	v.valueType = ValueShort
	binary.LittleEndian.PutUint16(v.valueBuffer[:], uint16(s))
}

func (v *EQValue) GetUShort() uint16 {
	return binary.LittleEndian.Uint16(v.valueBuffer[:])
}

func (v *EQValue) UShort(s uint16) {
	v.valueType = ValueUShort
	binary.LittleEndian.PutUint16(v.valueBuffer[:], s)
}

func (v *EQValue) GetInt() int {
	return int(binary.LittleEndian.Uint32(v.valueBuffer[:]))
}

func (v *EQValue) Int(i int) {
	v.valueType = ValueInt
	binary.LittleEndian.PutUint32(v.valueBuffer[:], uint32(i))
}

func (v *EQValue) GetUInt() uint {
	return uint(binary.LittleEndian.Uint32(v.valueBuffer[:]))
}

func (v *EQValue) UInt(u uint) {
	v.valueType = ValueUInt
	binary.LittleEndian.PutUint32(v.valueBuffer[:], uint32(u))
}

func (v *EQValue) GetInt64() int64 {
	return int64(binary.LittleEndian.Uint64(v.valueBuffer[:]))
}

func (v *EQValue) Int64(i int64) {
	v.valueType = ValueInt64
	binary.LittleEndian.PutUint64(v.valueBuffer[:], uint64(i))
}

func (v *EQValue) GetUInt64() uint64 {
	return binary.LittleEndian.Uint64(v.valueBuffer[:])
}

func (v *EQValue) UInt64(u uint64) {
	v.valueType = ValueUInt64
	binary.LittleEndian.PutUint64(v.valueBuffer[:], u)
}

func (v *EQValue) GetSingle() float32 {
	return math.Float32frombits(
		binary.LittleEndian.Uint32(v.valueBuffer[:]),
	)
}

func (v *EQValue) Single(s float32) {
	v.valueType = ValueSingle
	binary.LittleEndian.PutUint32(v.valueBuffer[:], math.Float32bits(s))
}

func (v *EQValue) GetDouble() float64 {
	return math.Float64frombits(
		binary.LittleEndian.Uint64(v.valueBuffer[:]),
	)
}

func (v *EQValue) Double(d float64) {
	v.valueType = ValueDouble
	binary.LittleEndian.PutUint64(v.valueBuffer[:], math.Float64bits(d))
}

func (v *EQValue) GetBytes() []byte {
	return v.valueBuffer[:]
}

func (v *EQValue) Bytes(d []byte) {
	v.valueType = ValueBytes
	copy(v.valueBuffer[:], d)
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
	Code       string
	Date       time.Time
	indicators []string
	value      []*EQValue

	extended map[string]*EQValue
}

func (v *Indicator) Extend(name string, value *EQValue) {
	if name == "" || value == nil {
		return
	}

	if v.extended == nil {
		v.extended = make(map[string]*EQValue)
	}
	v.extended[name] = value
}

func (v Indicator) GetExtend(name string) *EQValue {
	if v.extended == nil {
		return nil
	}

	return v.extended[name]
}

func (v Indicator) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("{Code:")
	buff.WriteString(v.Code)
	buff.WriteString(" Date:")
	buff.WriteString(v.Date.Format("2006-01-02"))
	buff.WriteString(" Indicators:{")
	for idx, name := range v.indicators {
		if idx > 0 {
			buff.WriteByte(' ')
		}
		fmt.Fprintf(buff, "%s:%+v", name, v.value[idx].GetValue())
	}
	buff.WriteString("}}")

	return buff.String()
}

func (v Indicator) Iter() func(yield func(string, *EQValue) bool) {
	return func(yield func(string, *EQValue) bool) {
		for idx, col := range v.indicators {
			if !yield(col, v.value[idx]) {
				return
			}
		}
	}
}

type EQData struct {
	codes      []string
	indicators []string
	dateList   []string
	values     []*EQValue
}

func (data *EQData) GetCodes() []string {
	return data.codes
}

func (data *EQData) GetDateList() []string {
	return data.dateList
}

func (data *EQData) GetIndicators() []string {
	return data.indicators
}

func (data *EQData) GetValues() []*EQValue {
	return data.values
}

func (data *EQData) Iter() func(yield func(int, *Indicator) bool) {
	codeSize := len(data.codes)
	indicatorSize := len(data.indicators)

	return func(yield func(int, *Indicator) bool) {
		rowIdx := 0

		for idxDate, dateStr := range data.dateList {
			var dateV []string
			switch {
			case strings.Contains(dateStr, "/"):
				dateV = strings.SplitN(dateStr, "/", 3)
			case strings.Contains(dateStr, "-"):
				dateV = strings.SplitN(dateStr, "-", 3)
			default:
				slog.Error(
					"unsupported date format",
					slog.String("date", dateStr),
				)
				return
			}

			if len(dateV) != 3 {
				slog.Error(
					"invalid date format",
					slog.String("date", dateStr),
				)
				return
			}

			yearV, err := strconv.Atoi(dateV[0])
			if err != nil {
				slog.Error(
					"parse date failed",
					slog.Any("error", err),
					slog.String("date", dateStr),
				)
				return
			}
			monthV, err := strconv.Atoi(dateV[1])
			if err != nil {
				slog.Error(
					"parse date failed",
					slog.Any("error", err),
					slog.String("date", dateStr),
				)
				return
			}
			dayV, err := strconv.Atoi(dateV[2])
			if err != nil {
				slog.Error(
					"parse date failed",
					slog.Any("error", err),
					slog.String("date", dateStr),
				)
				return
			}

			date := time.Date(
				yearV, time.Month(monthV), dayV,
				0, 0, 0, 0, time.Local,
			)

			for idxCode, code := range data.codes {
				value := Indicator{
					Code:       code,
					Date:       date,
					indicators: data.indicators,
					value:      make([]*EQValue, indicatorSize),
				}

				for idxIndicator := range data.indicators {
					idx := codeSize*indicatorSize*idxDate + indicatorSize*idxCode + idxIndicator
					value.value[idxIndicator] = data.values[idx]
				}

				if !yield(rowIdx, &value) {
					return
				}

				rowIdx++
			}
		}
	}
}

type Report struct {
	indicators []string
	value      []*EQValue

	extended map[string]*EQValue
}

func (rpt *Report) Extend(name string, value *EQValue) {
	if name == "" || value == nil {
		return
	}

	if rpt.extended == nil {
		rpt.extended = make(map[string]*EQValue)
	}
	rpt.extended[name] = value
}

func (rpt Report) GetExtend(name string) *EQValue {
	if rpt.extended == nil {
		return nil
	}

	return rpt.extended[name]
}

func (rpt Report) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("Report{")
	for idx, name := range rpt.indicators {
		if idx > 0 {
			buff.WriteByte(' ')
		}
		fmt.Fprintf(buff, "%s:%+v", name, rpt.value[idx].GetValue())
	}
	buff.WriteString("}")

	return buff.String()
}

func (rpt Report) Iter() func(yield func(string, *EQValue) bool) {
	return func(yield func(string, *EQValue) bool) {
		for idx, col := range rpt.indicators {
			if !yield(col, rpt.value[idx]) {
				return
			}
		}
	}
}

type EQCtrData struct {
	row        int
	column     int
	indicators []string
	values     []*EQValue
}

func (ctr *EQCtrData) Dimension() (int, int) {
	return ctr.row, ctr.column
}

func (ctr *EQCtrData) GetIndicators() []string {
	return ctr.indicators
}

func (ctr *EQCtrData) GetValues() []*EQValue {
	return ctr.values
}

func (ctr *EQCtrData) Iter() func(yield func(int, *Report) bool) {
	return func(yield func(int, *Report) bool) {
		for rowIdx := range ctr.row {
			value := Report{
				indicators: ctr.indicators,
				value:      make([]*EQValue, ctr.column),
			}

			for colIdx := range ctr.column {
				idx := ctr.column*rowIdx + colIdx
				value.value[colIdx] = ctr.values[idx]
			}

			if !yield(rowIdx, &value) {
				return
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
