package tpl

const Getter = `// Generated at {{.Time}}
package bean

type {{.Name}} struct {
	Calls   uint
	Returns {{.Type}}
}

func (g *{{.Name}}) Get() {{.Type}} {
	g.Calls++
	return g.Returns
}

func (g {{.Name}}) CallCount() uint {
	return g.Calls
}

func (g *{{.Name}}) SetReturnValue(val {{.Type}}) {
	g.Returns = val
}
`
