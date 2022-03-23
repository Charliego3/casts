package casts

import (
	"github.com/spf13/cast"
	"time"
)

func MustBool(i any) bool {
	return cast.ToBool(i)
}

func Bool(i any) (bool, error) {
	return cast.ToBoolE(i)
}

func BoolDefault(i any, dv bool) bool {
	v, err := cast.ToBoolE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustDuration(i any) time.Duration {
	return cast.ToDuration(i)
}

func Duration(i any) (time.Duration, error) {
	return cast.ToDurationE(i)
}

func DurationDefault(i any, dv time.Duration) time.Duration {
	v, err := cast.ToDurationE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustFloat32(i any) float32 {
	return cast.ToFloat32(i)
}

func Float32(i any) (float32, error) {
	return cast.ToFloat32E(i)
}

func Float32Default(i any, dv float32) float32 {
	v, err := cast.ToFloat32E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustFloat64(i any) float64 {
	return cast.ToFloat64(i)
}

func Float64(i any) (float64, error) {
	return cast.ToFloat64E(i)
}

func Float64Default(i any, dv float64) float64 {
	v, err := cast.ToFloat64E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustInt(i any) int {
	return cast.ToInt(i)
}

func Int(i any) (int, error) {
	return cast.ToIntE(i)
}

func IntDefault(i any, dv int) int {
	v, err := cast.ToIntE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustInt8(i any) int8 {
	return cast.ToInt8(i)
}

func Int8(i any) (int8, error) {
	return cast.ToInt8E(i)
}

func Int8Default(i any, dv int8) int8 {
	v, err := cast.ToInt8E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustInt16(i any) int16 {
	return cast.ToInt16(i)
}

func Int16(i any) (int16, error) {
	return cast.ToInt16E(i)
}

func Int16Default(i any, dv int16) int16 {
	v, err := cast.ToInt16E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustInt32(i any) int32 {
	return cast.ToInt32(i)
}

func Int32(i any) (int32, error) {
	return cast.ToInt32E(i)
}

func Int32Default(i any, dv int32) int32 {
	v, err := cast.ToInt32E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustInt64(i any) int64 {
	return cast.ToInt64(i)
}

func Int64(i any) (int64, error) {
	return cast.ToInt64E(i)
}

func Int64Default(i any, dv int64) int64 {
	v, err := cast.ToInt64E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustString(i any) string {
	return cast.ToString(i)
}

func String(i any) (string, error) {
	return cast.ToStringE(i)
}

func StringDefault(i any, dv string) string {
	v, err := cast.ToStringE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustUint(i any) uint {
	return cast.ToUint(i)
}

func Uint(i any) (uint, error) {
	return cast.ToUintE(i)
}

func UintDefault(i any, dv uint) uint {
	v, err := cast.ToUintE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustUint8(i any) uint8 {
	return cast.ToUint8(i)
}

func Uint8(i any) (uint8, error) {
	return cast.ToUint8E(i)
}

func Uint8Default(i any, dv uint8) uint8 {
	v, err := cast.ToUint8E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustUint16(i any) uint16 {
	return cast.ToUint16(i)
}

func Uint16(i any) (uint16, error) {
	return cast.ToUint16E(i)
}

func Uint16Default(i any, dv uint16) uint16 {
	v, err := cast.ToUint16E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustUint32(i any) uint32 {
	return cast.ToUint32(i)
}

func Uint32(i any) (uint32, error) {
	return cast.ToUint32E(i)
}

func Uint32Default(i any, dv uint32) uint32 {
	v, err := cast.ToUint32E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustUint64(i any) uint64 {
	return cast.ToUint64(i)
}

func Uint64(i any) (uint64, error) {
	return cast.ToUint64E(i)
}

func Uint64Default(i any, dv uint64) uint64 {
	v, err := cast.ToUint64E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustTime(i any) time.Time {
	return cast.ToTime(i)
}

func Time(i any) (time.Time, error) {
	return cast.ToTimeE(i)
}

func TimeDefault(i any, dv time.Time) time.Time {
	v, err := cast.ToTimeE(i)
	if err != nil {
		return dv
	}
	return v
}
