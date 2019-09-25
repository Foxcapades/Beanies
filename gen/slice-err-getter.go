package main

import (
	. "github.com/Foxcapades/Beanies/gen/tpl"
	. "github.com/Foxcapades/Beanies/gen/util"
	. "text/template"
)

func main() {
	Execute(ParseInput(sliceErrGetMap), Must(New("getter").Parse(ErrGetter)))
}

func sliceErrGetMap(def TypeDef) TypeDef {
	def.Name = def.Name + SuffixSliceErrGetter
	def.Type = "[]" + def.Type
	return def
}
