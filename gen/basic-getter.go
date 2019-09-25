package main

import (
	. "github.com/Foxcapades/Beanies/gen/tpl"
	. "github.com/Foxcapades/Beanies/gen/util"
	. "text/template"
)

func main() {
	Execute(ParseInput(basicTransform), Must(New("getter").Parse(Getter)))
}

func basicTransform(def TypeDef) TypeDef {
	def.Name = def.Name + SuffixGetter
	return def
}
