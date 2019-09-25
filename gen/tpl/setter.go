package tpl

const Setter = `// Generated at {{.Time}}
package bean

type {{.Name}} struct {
	Calls   uint
	Inputs  []{{.Type}}
}

func (g *{{.Name}}) Set(v {{.Type}}) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g {{.Name}}) CallCount() uint {
	return g.Calls
}

func (g *{{.Name}}) InputValues() []{{.Type}} {
	return g.Inputs
}
`
