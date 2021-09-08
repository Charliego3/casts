package main

import (
	"fmt"
	"os"
	"strings"
)

type cast struct {
	alias       string
	types       string
	slice       bool
	isMap       bool
	onlySlice   bool
	appendSlice bool
}

type imp struct {
	url   string
	slice bool
	maps  bool
	cast  bool
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
	mustMapTemplate = `func StringMap%s(i interface{}) map[string]%s {
	return cast.ToStringMap%s(i)
}%s
`

	defaultMapTemplate = `func StringMap%sDefault(i interface{}, di map[string]%s) map[string]%s {
	val, err := cast.ToStringMap%sE(i)
	if err != nil {
		return di
	}
	return val
}%s
`

	castFileName  = "cast.go"
	castSliceName = "cast_slice.go"
	castMapName   = "cast_map.go"
)

var (
	files   = make(map[string]*os.File)
	imports = []imp{
		{url: "github.com/spf13/cast", slice: true, maps: true, cast: true},
		{url: "time", slice: true, maps: false, cast: true},
	}
)

func main() {
	allType := []cast{
		{"bool", "bool", true, false, false, true},
		{"duration", "time.Duration", true, false, false, true},
		{"float32", "float32", false, false, false, true},
		{"float64", "float64", false, false, false, true},
		{"int16", "int16", false, false, false, true},
		{"int32", "int32", false, false, false, true},
		{"int64", "int64", false, false, false, true},
		{"int8", "int8", false, false, false, true},
		{"int", "int", true, false, false, true},
		{"slice", "interface{}", true, false, true, false},
		{"string", "string", true, false, false, true},
		{"time", "time.Time", false, false, false, false},
		{"uint", "uint", false, false, false, false},
		{"uint16", "uint16", false, false, false, false},
		{"uint32", "uint32", false, false, false, false},
		{"uint64", "uint64", false, false, false, false},
		{"uint8", "uint8", false, false, false, false},

		{"", "interface{}", false, true, false, false},
		{"bool", "bool", false, true, false, false},
		{"int", "int", false, true, false, false},
		{"int64", "int64", false, true, false, false},
		{"string", "string", false, true, false, false},
		{"stringSlice", "[]string", false, true, false, false},
	}

	execute(allType)
}

func execute(casts []cast) {
	castFile := getFile(castFileName)
	sliceFile := getFile(castSliceName)
	mapFile := getFile(castMapName)
	defer func() {
		_ = castFile.Close()
		_ = sliceFile.Close()
		_ = mapFile.Close()
	}()

	writeHeader(castFile)
	writeHeader(sliceFile)
	writeHeader(mapFile)

	for _, c := range casts {
		var title, slice, types, line string
		title = strings.Title(c.alias)
		types = c.types
		line = "\n"

		if c.isMap {
			writeFile(mapFile, mustMapTemplate, title, types, title, line)
			writeFile(mapFile, defaultMapTemplate, title, types, types, title, line)
			continue
		}

		if !c.onlySlice {
			writeFile(castFile, mustTemplate, title, slice, types, title, slice, line)
			writeFile(castFile, typeTemplate, title, slice, types, title, slice, line)
			writeFile(castFile, defaultTemplate, title, slice, types, types, title, slice, line)
		}

		if c.slice {
			types = "[]" + types
			if c.appendSlice {
				slice = "Slice"
			}
			writeFile(sliceFile, mustTemplate, title, slice, types, title, slice, line)
			writeFile(sliceFile, typeTemplate, title, slice, types, title, slice, line)
			writeFile(sliceFile, defaultTemplate, title, slice, types, types, title, slice, line)
		}
	}
}

func writeFile(file *os.File, format string, args ...interface{}) {
	_, err := fmt.Fprintf(file, format, args...)
	if err != nil {
		panic(err)
	}
}

func writeHeader(file *os.File) {
	headers := "package casts \n\n"
	headers += "import (\n"

	println(file.Name())

	for _, s := range imports {
		write := file.Name() == castFileName && s.cast
		write = (file.Name() == castSliceName && s.slice) || write
		write = (file.Name() == castMapName && s.maps) || write

		if write {
			headers += fmt.Sprintf("\t%q\n", s.url)
		}
	}

	headers += ")\n\n"
	_, err := fmt.Fprintf(file, headers)
	if err != nil {
		panic(err)
	}
}

func getFile(filename string) *os.File {
	if file, ok := files[filename]; ok {
		return file
	}
	file := openFile(filename)
	files[filename] = file
	return file
}

func openFile(fn string) *os.File {
	file, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return file
}
