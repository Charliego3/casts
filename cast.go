package casts 

import (
	"github.com/spf13/cast"
	"time"
)

func MustBool(i interface{}) bool {
	return cast.ToBool(i)
}

func Bool(i interface{}) (bool, error) {
	return cast.ToBoolE(i)
}

func BoolDefault(i interface{}, di bool) bool {
	val, err := cast.ToBoolE(i)
	if err != nil {
		return di
	}
	return val
}

func MustDuration(i interface{}) time.Duration {
	return cast.ToDuration(i)
}

func Duration(i interface{}) (time.Duration, error) {
	return cast.ToDurationE(i)
}

func DurationDefault(i interface{}, di time.Duration) time.Duration {
	val, err := cast.ToDurationE(i)
	if err != nil {
		return di
	}
	return val
}

func MustFloat32(i interface{}) float32 {
	return cast.ToFloat32(i)
}

func Float32(i interface{}) (float32, error) {
	return cast.ToFloat32E(i)
}

func Float32Default(i interface{}, di float32) float32 {
	val, err := cast.ToFloat32E(i)
	if err != nil {
		return di
	}
	return val
}

func MustFloat64(i interface{}) float64 {
	return cast.ToFloat64(i)
}

func Float64(i interface{}) (float64, error) {
	return cast.ToFloat64E(i)
}

func Float64Default(i interface{}, di float64) float64 {
	val, err := cast.ToFloat64E(i)
	if err != nil {
		return di
	}
	return val
}

func MustInt16(i interface{}) int16 {
	return cast.ToInt16(i)
}

func Int16(i interface{}) (int16, error) {
	return cast.ToInt16E(i)
}

func Int16Default(i interface{}, di int16) int16 {
	val, err := cast.ToInt16E(i)
	if err != nil {
		return di
	}
	return val
}

func MustInt32(i interface{}) int32 {
	return cast.ToInt32(i)
}

func Int32(i interface{}) (int32, error) {
	return cast.ToInt32E(i)
}

func Int32Default(i interface{}, di int32) int32 {
	val, err := cast.ToInt32E(i)
	if err != nil {
		return di
	}
	return val
}

func MustInt64(i interface{}) int64 {
	return cast.ToInt64(i)
}

func Int64(i interface{}) (int64, error) {
	return cast.ToInt64E(i)
}

func Int64Default(i interface{}, di int64) int64 {
	val, err := cast.ToInt64E(i)
	if err != nil {
		return di
	}
	return val
}

func MustInt8(i interface{}) int8 {
	return cast.ToInt8(i)
}

func Int8(i interface{}) (int8, error) {
	return cast.ToInt8E(i)
}

func Int8Default(i interface{}, di int8) int8 {
	val, err := cast.ToInt8E(i)
	if err != nil {
		return di
	}
	return val
}

func MustInt(i interface{}) int {
	return cast.ToInt(i)
}

func Int(i interface{}) (int, error) {
	return cast.ToIntE(i)
}

func IntDefault(i interface{}, di int) int {
	val, err := cast.ToIntE(i)
	if err != nil {
		return di
	}
	return val
}

func MustString(i interface{}) string {
	return cast.ToString(i)
}

func String(i interface{}) (string, error) {
	return cast.ToStringE(i)
}

func StringDefault(i interface{}, di string) string {
	val, err := cast.ToStringE(i)
	if err != nil {
		return di
	}
	return val
}

func MustTime(i interface{}) time.Time {
	return cast.ToTime(i)
}

func Time(i interface{}) (time.Time, error) {
	return cast.ToTimeE(i)
}

func TimeDefault(i interface{}, di time.Time) time.Time {
	val, err := cast.ToTimeE(i)
	if err != nil {
		return di
	}
	return val
}

func MustUint(i interface{}) uint {
	return cast.ToUint(i)
}

func Uint(i interface{}) (uint, error) {
	return cast.ToUintE(i)
}

func UintDefault(i interface{}, di uint) uint {
	val, err := cast.ToUintE(i)
	if err != nil {
		return di
	}
	return val
}

func MustUint16(i interface{}) uint16 {
	return cast.ToUint16(i)
}

func Uint16(i interface{}) (uint16, error) {
	return cast.ToUint16E(i)
}

func Uint16Default(i interface{}, di uint16) uint16 {
	val, err := cast.ToUint16E(i)
	if err != nil {
		return di
	}
	return val
}

func MustUint32(i interface{}) uint32 {
	return cast.ToUint32(i)
}

func Uint32(i interface{}) (uint32, error) {
	return cast.ToUint32E(i)
}

func Uint32Default(i interface{}, di uint32) uint32 {
	val, err := cast.ToUint32E(i)
	if err != nil {
		return di
	}
	return val
}

func MustUint64(i interface{}) uint64 {
	return cast.ToUint64(i)
}

func Uint64(i interface{}) (uint64, error) {
	return cast.ToUint64E(i)
}

func Uint64Default(i interface{}, di uint64) uint64 {
	val, err := cast.ToUint64E(i)
	if err != nil {
		return di
	}
	return val
}

func MustUint8(i interface{}) uint8 {
	return cast.ToUint8(i)
}

func Uint8(i interface{}) (uint8, error) {
	return cast.ToUint8E(i)
}

func Uint8Default(i interface{}, di uint8) uint8 {
	val, err := cast.ToUint8E(i)
	if err != nil {
		return di
	}
	return val
}

