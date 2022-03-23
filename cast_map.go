package casts

import "github.com/spf13/cast"

func MustMap(i any) map[string]any {
	return cast.ToStringMap(i)
}

func Map(i any) (map[string]any, error) {
	return cast.ToStringMapE(i)
}

func MapDefault(i any, dv map[string]any) map[string]any {
	v, err := cast.ToStringMapE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustMapBool(i any) map[string]bool {
	return cast.ToStringMapBool(i)
}

func MapBool(i any) (map[string]bool, error) {
	return cast.ToStringMapBoolE(i)
}

func MapBoolDefault(i any, dv map[string]bool) map[string]bool {
	v, err := cast.ToStringMapBoolE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustMapInt(i any) map[string]int {
	return cast.ToStringMapInt(i)
}

func MapInt(i any) (map[string]int, error) {
	return cast.ToStringMapIntE(i)
}

func MapIntDefault(i any, dv map[string]int) map[string]int {
	v, err := cast.ToStringMapIntE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustMapInt64(i any) map[string]int64 {
	return cast.ToStringMapInt64(i)
}

func MapInt64(i any) (map[string]int64, error) {
	return cast.ToStringMapInt64E(i)
}

func MapInt64Default(i any, dv map[string]int64) map[string]int64 {
	v, err := cast.ToStringMapInt64E(i)
	if err != nil {
		return dv
	}
	return v
}

func MustMapString(i any) map[string]string {
	return cast.ToStringMapString(i)
}

func MapString(i any) (map[string]string, error) {
	return cast.ToStringMapStringE(i)
}

func MapStringDefault(i any, dv map[string]string) map[string]string {
	v, err := cast.ToStringMapStringE(i)
	if err != nil {
		return dv
	}
	return v
}

func MustMapStringSlice(i any) map[string][]string {
	return cast.ToStringMapStringSlice(i)
}

func MapStringSlice(i any) (map[string][]string, error) {
	return cast.ToStringMapStringSliceE(i)
}

func MapStringSliceDefault(i any, dv map[string][]string) map[string][]string {
	v, err := cast.ToStringMapStringSliceE(i)
	if err != nil {
		return dv
	}
	return v
}
