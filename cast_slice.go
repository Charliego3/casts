package casts

import (
	"github.com/spf13/cast"
	"time"
)

func MustSlice(i any) []any {
	return cast.ToSlice(i)
}

func Slice(i any) ([]any, error) {
	return cast.ToSliceE(i)
}

func SliceDefault(i any, dv []any) []any {
	v, err := cast.ToSliceE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustSliceInt(i any) []int {
	return cast.ToIntSlice(i)
}

func SliceInt(i any) ([]int, error) {
	return cast.ToIntSliceE(i)
}

func SliceIntDefault(i any, dv []int) []int {
	v, err := cast.ToIntSliceE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustSliceBool(i any) []bool {
	return cast.ToBoolSlice(i)
}

func SliceBool(i any) ([]bool, error) {
	return cast.ToBoolSliceE(i)
}

func SliceBoolDefault(i any, dv []bool) []bool {
	v, err := cast.ToBoolSliceE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustSliceString(i any) []string {
	return cast.ToStringSlice(i)
}

func SliceString(i any) ([]string, error) {
	return cast.ToStringSliceE(i)
}

func SliceStringDefault(i any, dv []string) []string {
	v, err := cast.ToStringSliceE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustSliceDuration(i any) []time.Duration {
	return cast.ToDurationSlice(i)
}

func SliceDuration(i any) ([]time.Duration, error) {
	return cast.ToDurationSliceE(i)
}

func SliceDurationDefault(i any, dv []time.Duration) []time.Duration {
	v, err := cast.ToDurationSliceE(i)
	if err != nil {
		return dv
	}
	return v
}
