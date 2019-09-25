package tpl

const ErrSetter = `// Generated at {{.Time}}
package bean

type {{.Name}} struct {
	Calls   uint
	Inputs  []{{.Type}}
	Error   error
}

func (g *{{.Name}}) Set(v {{.Type}}) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g {{.Name}}) CallCount() uint {
	return g.Calls
}

func (g *{{.Name}}) InputValues() []{{.Type}} {
	return g.Inputs
}

func (g *{{.Name}}) SetError(err error) {
	g.Error = err
}
`
