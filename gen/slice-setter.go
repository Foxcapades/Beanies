package main

import (
	. "github.com/Foxcapades/Beanies/gen/tpl"
	. "github.com/Foxcapades/Beanies/gen/util"
	. "text/template"
)

func main() {
	Execute(ParseInput(sliceSetMap), Must(New("setter").Parse(Setter)))
}

func sliceSetMap(def TypeDef) TypeDef {
	def.Name = def.Name + SuffixSliceSetter
	def.Type = "[]" + def.Type
	return def
}
