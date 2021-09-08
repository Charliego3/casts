package main

import (
	"fmt"
	"os"
	"strings"
)

type cast struct {
	alias string
	types string
	slice bool
	onlySlice bool
	appendSlice bool
}

const (
	mustTemplate = `func Must%s%s(i interface{}) %s {
	return cast.To%s%s(i)
}%s
`
	typeTemplate = `func %s%s(i interface{}) (%s, error) {
	return cast.To%s%sE(i)
}%s
`
	defaultTemplate = `func %s%sDefault(i interface{}, di %s) %s {
	val, err := cast.To%s%sE(i)
	if err != nil {
		return di
	}
	return val
}%s
`
)

func main() {
	imports := []string{
		"github.com/spf13/cast",
		"time",
	}

	allType := []cast{
		{"bool", "bool", true, false, true},
		{"duration", "time.Duration", true, false, true},
		{"float32", "float32", false, false, true},
		{"float64", "float64", false, false, true},
		{"int16", "int16", false, false, true},
		{"int32", "int32", false, false, true},
		{"int64", "int64", false, false, true},
		{"int8", "int8", false, false, true},
		{"int", "int", true, false, true},
		{"slice", "interface{}", true, true, false},
		{"string", "string", true, false, true},
		{"time", "time.Time", false, false, false},
		{"uint", "uint", false, false, false},
		{"uint16", "uint16", false, false, false},
		{"uint32", "uint32", false, false, false},
		{"uint64", "uint64", false, false, false},
		{"uint8", "uint8", false, false, false},
	}

	castFunc(imports, allType, false)
	castFunc(imports, allType, true)
}

func castFunc(imports []string, casts []cast, slice bool) {
	fn := "cast.go"
	if slice {
		fn = "cast_slice.go"
	}
	file := openFile(fn)
	defer file.Close()

	headers := "package casts \n\nimport (\n"
	for _, s := range imports {
		headers += fmt.Sprintf("\t%q\n", s)
	}
	headers += ")\n\n"
	fmt.Fprintf(file, headers)


	line := "\n"
	i := 0
	for _, c := range casts {
		sr := ""
		t := c.types
		if slice {
			if !c.slice {
				continue
			} else {
				t = "[]"+t
			}
			if c.appendSlice {
				fmt.Printf("%#v\n", c)
				sr = "Slice"
			}
		} else if c.onlySlice {
			continue
		}

		title := strings.Title(c.alias)
		fmt.Fprintf(file, mustTemplate, title, sr, t, title, sr, line)
		fmt.Fprintf(file, typeTemplate, title, sr, t, title, sr, line)

		if i == len(casts)-1 {
			line = ""
		}
		fmt.Fprintf(file, defaultTemplate, title, sr, t, t, title, sr, line)
	}
}

func openFile(fn string) *os.File {
	file, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return file
}
