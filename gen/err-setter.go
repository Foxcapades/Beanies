package main

import (
	. "github.com/Foxcapades/Beanies/gen/tpl"
	. "github.com/Foxcapades/Beanies/gen/util"
	. "text/template"
)

func main() {
	Execute(ParseInput(errSetMap), Must(New("setter").Parse(ErrSetter)))
}

func errSetMap(def TypeDef) TypeDef {
	def.Name = def.Name + SuffixErrorSetter
	return def
}
