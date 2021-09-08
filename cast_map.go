package casts

import (
	"github.com/spf13/cast"
)

func StringMap(i interface{}) map[string]interface{} {
	return cast.ToStringMap(i)
}

func StringMapDefault(i interface{}, di map[string]interface{}) map[string]interface{} {
	val, err := cast.ToStringMapE(i)
	if err != nil {
		return di
	}
	return val
}

func StringMapBool(i interface{}) map[string]bool {
	return cast.ToStringMapBool(i)
}

func StringMapBoolDefault(i interface{}, di map[string]bool) map[string]bool {
	val, err := cast.ToStringMapBoolE(i)
	if err != nil {
		return di
	}
	return val
}

func StringMapInt(i interface{}) map[string]int {
	return cast.ToStringMapInt(i)
}

func StringMapIntDefault(i interface{}, di map[string]int) map[string]int {
	val, err := cast.ToStringMapIntE(i)
	if err != nil {
		return di
	}
	return val
}

func StringMapInt64(i interface{}) map[string]int64 {
	return cast.ToStringMapInt64(i)
}

func StringMapInt64Default(i interface{}, di map[string]int64) map[string]int64 {
	val, err := cast.ToStringMapInt64E(i)
	if err != nil {
		return di
	}
	return val
}

func StringMapString(i interface{}) map[string]string {
	return cast.ToStringMapString(i)
}

func StringMapStringDefault(i interface{}, di map[string]string) map[string]string {
	val, err := cast.ToStringMapStringE(i)
	if err != nil {
		return di
	}
	return val
}

func StringMapStringSlice(i interface{}) map[string][]string {
	return cast.ToStringMapStringSlice(i)
}

func StringMapStringSliceDefault(i interface{}, di map[string][]string) map[string][]string {
	val, err := cast.ToStringMapStringSliceE(i)
	if err != nil {
		return di
	}
	return val
}
