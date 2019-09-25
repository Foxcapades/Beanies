package main

import (
	. "github.com/Foxcapades/Beanies/gen/tpl"
	. "github.com/Foxcapades/Beanies/gen/util"
	. "text/template"
)

func main() {
	Execute(ParseInput(simpleErrGetMap), Must(New("getter").Parse(ErrGetter)))
}

func simpleErrGetMap(def TypeDef) TypeDef {
	def.Name = def.Name + SuffixErrorGetter
	return def
}
