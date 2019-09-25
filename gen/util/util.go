package util

import (
	"flag"
	"fmt"
	"os"
	"text/template"
	"time"
)

type TypeDef struct {
	Time string
	Name string
	Type string
}

const (
	SuffixGetter         = "Getter"
	SuffixSetter         = "Setter"
	SuffixErrorGetter    = "Err" + SuffixGetter
	SuffixErrorSetter    = "Err" + SuffixSetter
	SuffixSliceGetter    = "Slice" + SuffixGetter
	SuffixSliceSetter    = "Slice" + SuffixSetter
	SuffixSliceErrGetter = "Slice" + SuffixErrorGetter
	SuffixSliceErrSetter = "Slice" + SuffixErrorSetter
)

type Mapper = func(TypeDef) TypeDef

func ParseInput(nMap Mapper) TypeDef {
	var data TypeDef
	flag.StringVar(&data.Type, "type", "", "Go type name for the generated mock bean")
	flag.StringVar(&data.Name, "name", "", "Name prefix for the generated mock bean")
	flag.Parse()

	if data.Name == "" {
		_, _ = fmt.Fprint(os.Stderr, "-name is required")
		os.Exit(1)
	}
	if data.Type == "" {
		_, _ = fmt.Fprint(os.Stderr, "-type is required")
		os.Exit(1)
	}

	data.Time = time.Now().Format(time.RFC3339)

	return nMap(data)
}

func Execute(def TypeDef, tpl *template.Template) {
	if e := tpl.Execute(os.Stdout, def); e != nil {
		panic(e)
	}
}


