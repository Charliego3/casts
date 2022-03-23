package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	context        = "return cast.To%s(i)"
	contextE       = "return cast.To%sE(i)"
	contextDefault = `v, err := cast.To%sE(i)
	if err != nil {
		return dv
	}
	return v`

	fnt = `func %s(i any%s) %s {
	%s
}`
)

type Template struct {
	f       *os.File
	imports []string
	fns     []fn
	prefix  string
}

type fn struct {
	name    string
	cast    string
	types   string
	options []option
}

type option struct {
	prefix  string
	suffix  string
	rt      string
	context string
}

func main() {
	dos := []option{
		{"Must", "", "", context},
		{"", "", "(%s, error)", contextE},
		{"", "Default", "", contextDefault},
	}

	ts := []Template{
		{f: openFile("cast.go"), imports: []string{
			"github.com/spf13/cast",
			"time",
		}, fns: []fn{
			{"Bool", "Bool", "bool", dos},
			{"Duration", "Duration", "time.Duration", dos},
			{"Float32", "Float32", "float32", dos},
			{"Float64", "Float64", "float64", dos},
			{"Int", "Int", "int", dos},
			{"Int8", "Int8", "int8", dos},
			{"Int16", "Int16", "int16", dos},
			{"Int32", "Int32", "int32", dos},
			{"Int64", "Int64", "int64", dos},
			{"String", "String", "string", dos},
			{"Uint", "Uint", "uint", dos},
			{"Uint8", "Uint8", "uint8", dos},
			{"Uint16", "Uint16", "uint16", dos},
			{"Uint32", "Uint32", "uint32", dos},
			{"Uint64", "Uint64", "uint64", dos},
			{"Time", "Time", "time.Time", dos},
		}},

		{f: openFile("cast_slice.go"), imports: []string{
			"github.com/spf13/cast",
			"time",
		}, fns: []fn{
			{"Slice", "Slice", "[]any", dos},
			{"SliceInt", "IntSlice", "[]int", dos},
			{"SliceBool", "BoolSlice", "[]bool", dos},
			{"SliceString", "StringSlice", "[]string", dos},
			{"SliceDuration", "DurationSlice", "[]time.Duration", dos},
		}},

		{f: openFile("cast_map.go"), imports: []string{
			"github.com/spf13/cast",
		}, prefix: "StringMap", fns: []fn{
			{"Map", "", "map[string]any", dos},
			{"MapBool", "Bool", "map[string]bool", dos},
			{"MapInt", "Int", "map[string]int", dos},
			{"MapInt64", "Int64", "map[string]int64", dos},
			{"MapString", "String", "map[string]string", dos},
			{"MapStringSlice", "StringSlice", "map[string][]string", dos},
		}},
	}

	ph := "package casts\n\n"
	for _, t := range ts {
		content := bufio.NewWriter(t.f)
		_, _ = content.WriteString(ph)

		il := len(t.imports)
		if il > 0 {
			_, _ = content.WriteString("import ")
		}
		if il > 1 {
			_, _ = content.WriteString("(\n")
		}

		for i, im := range t.imports {
			if i > 0 {
				_ = content.WriteByte('\n')
			}
			_, _ = content.WriteString(fmt.Sprintf("\t%q", im))
		}

		if il > 1 {
			_, _ = content.WriteString("\n)")
		}

		for _, f := range t.fns {
			for _, o := range f.options {
				_, _ = content.WriteString("\n\n")
				name := o.prefix + f.name + o.suffix
				df := ""
				if o.suffix == "Default" {
					df = ", dv " + f.types
				}
				rt := f.types
				if len(o.rt) > 0 {
					rt = fmt.Sprintf(o.rt, f.types)
				}

				c := fmt.Sprintf(o.context, t.prefix+f.cast)
				_, _ = content.WriteString(fmt.Sprintf(fnt, name, df, rt, c))
			}
		}

		err := content.Flush()
		if err != nil {
			panic(err)
		}
	}
}

func openFile(fn string) *os.File {
	file, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return file
}
