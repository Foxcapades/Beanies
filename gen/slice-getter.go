package main

import (
	. "github.com/Foxcapades/Beanies/gen/tpl"
	. "github.com/Foxcapades/Beanies/gen/util"
	. "text/template"
)

func main() {
	Execute(ParseInput(sliceTransform), Must(New("getter").Parse(Getter)))
}

func sliceTransform(def TypeDef) TypeDef {
	def.Name = def.Name + SuffixSliceGetter
	def.Type = "[]" + def.Type
	return def
}
