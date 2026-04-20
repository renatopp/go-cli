package cli

import (
	"time"

	v2 "github.com/renatopp/go-cli/v2"
)

func _addpos[T v2.Positional](a T) T {
	app.CurrentCommand().WithPositional(a)
	return a
}

func Pos(name, description string) *v2.GenericPositional[string] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseString))
}
func PosInt(name, description string) *v2.GenericPositional[int] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseInt[int]))
}
func PosInt8(name, description string) *v2.GenericPositional[int8] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseInt[int8]))
}
func PosInt16(name, description string) *v2.GenericPositional[int16] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseInt[int16]))
}
func PosInt32(name, description string) *v2.GenericPositional[int32] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseInt[int32]))
}
func PosInt64(name, description string) *v2.GenericPositional[int64] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseInt[int64]))
}
func PosUint(name, description string) *v2.GenericPositional[uint] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseUint[uint]))
}
func PosUint8(name, description string) *v2.GenericPositional[uint8] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseUint[uint8]))
}
func PosUint16(name, description string) *v2.GenericPositional[uint16] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseUint[uint16]))
}
func PosUint32(name, description string) *v2.GenericPositional[uint32] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseUint[uint32]))
}
func PosUint64(name, description string) *v2.GenericPositional[uint64] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseUint[uint64]))
}
func PosFloat32(name, description string) *v2.GenericPositional[float32] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseFloat[float32]))
}
func PosFloat64(name, description string) *v2.GenericPositional[float64] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseFloat[float64]))
}
func PosBool(name, description string) *v2.GenericPositional[bool] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseBool))
}
func PosDuration(name, description string) *v2.GenericPositional[time.Duration] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseDuration))
}

func PosIntSlice(name, description string) *v2.GenericPositional[[]int] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseIntSlice[int]))
}
func PosInt8Slice(name, description string) *v2.GenericPositional[[]int8] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseIntSlice[int8]))
}
func PosInt16Slice(name, description string) *v2.GenericPositional[[]int16] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseIntSlice[int16]))
}
func PosInt32Slice(name, description string) *v2.GenericPositional[[]int32] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseIntSlice[int32]))
}
func PosInt64Slice(name, description string) *v2.GenericPositional[[]int64] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseIntSlice[int64]))
}
func PosUintSlice(name, description string) *v2.GenericPositional[[]uint] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseUintSlice[uint]))
}
func PosUint8Slice(name, description string) *v2.GenericPositional[[]uint8] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseUintSlice[uint8]))
}
func PosUint16Slice(name, description string) *v2.GenericPositional[[]uint16] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseUintSlice[uint16]))
}
func PosUint32Slice(name, description string) *v2.GenericPositional[[]uint32] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseUintSlice[uint32]))
}
func PosUint64Slice(name, description string) *v2.GenericPositional[[]uint64] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseUintSlice[uint64]))
}
func PosFloat32Slice(name, description string) *v2.GenericPositional[[]float32] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseFloatSlice[float32]))
}
func PosFloat64Slice(name, description string) *v2.GenericPositional[[]float64] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseFloatSlice[float64]))
}
func PosBoolSlice(name, description string) *v2.GenericPositional[[]bool] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseBoolSlice))
}
func PosDurationSlice(name, description string) *v2.GenericPositional[[]time.Duration] {
	return _addpos(v2.NewGenericPositional(name, description, v2.ParseDurationSlice))
}
