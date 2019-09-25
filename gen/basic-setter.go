package main

import (
	. "github.com/Foxcapades/Beanies/gen/tpl"
	. "github.com/Foxcapades/Beanies/gen/util"
	. "text/template"
)

func main() {
	Execute(ParseInput(simpleSetMap), Must(New("setter").Parse(Setter)))
}

func simpleSetMap(def TypeDef) TypeDef {
	def.Name = def.Name + SuffixSetter
	return def
}
