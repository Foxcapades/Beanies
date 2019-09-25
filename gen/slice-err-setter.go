package main

import (
	. "github.com/Foxcapades/Beanies/gen/tpl"
	. "github.com/Foxcapades/Beanies/gen/util"
	. "text/template"
)

func main() {
	Execute(ParseInput(sliceErrSetMap), Must(New("setter").Parse(Setter)))
}

func sliceErrSetMap(def TypeDef) TypeDef {
	def.Name = def.Name + SuffixSliceErrSetter
	def.Type = "[]" + def.Type
	return def
}
