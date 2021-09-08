package casts

import (
	"github.com/spf13/cast"
	"time"
)

func MustBoolSlice(i interface{}) []bool {
	return cast.ToBoolSlice(i)
}

func BoolSlice(i interface{}) ([]bool, error) {
	return cast.ToBoolSliceE(i)
}

func BoolSliceDefault(i interface{}, di []bool) []bool {
	val, err := cast.ToBoolSliceE(i)
	if err != nil {
		return di
	}
	return val
}

func MustDurationSlice(i interface{}) []time.Duration {
	return cast.ToDurationSlice(i)
}

func DurationSlice(i interface{}) ([]time.Duration, error) {
	return cast.ToDurationSliceE(i)
}

func DurationSliceDefault(i interface{}, di []time.Duration) []time.Duration {
	val, err := cast.ToDurationSliceE(i)
	if err != nil {
		return di
	}
	return val
}

func MustIntSlice(i interface{}) []int {
	return cast.ToIntSlice(i)
}

func IntSlice(i interface{}) ([]int, error) {
	return cast.ToIntSliceE(i)
}

func IntSliceDefault(i interface{}, di []int) []int {
	val, err := cast.ToIntSliceE(i)
	if err != nil {
		return di
	}
	return val
}

func MustSlice(i interface{}) []interface{} {
	return cast.ToSlice(i)
}

func Slice(i interface{}) ([]interface{}, error) {
	return cast.ToSliceE(i)
}

func SliceDefault(i interface{}, di []interface{}) []interface{} {
	val, err := cast.ToSliceE(i)
	if err != nil {
		return di
	}
	return val
}

func MustStringSlice(i interface{}) []string {
	return cast.ToStringSlice(i)
}

func StringSlice(i interface{}) ([]string, error) {
	return cast.ToStringSliceE(i)
}

func StringSliceDefault(i interface{}, di []string) []string {
	val, err := cast.ToStringSliceE(i)
	if err != nil {
		return di
	}
	return val
}
