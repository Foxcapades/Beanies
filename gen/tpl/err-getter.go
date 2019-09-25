package tpl

const ErrGetter = `// Generated at {{.Time}}
package bean

type {{.Name}} struct {
	Calls   uint
	Returns {{.Type}}
	Error		error
}

func (g *{{.Name}}) Get() ({{.Type}}, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g {{.Name}}) CallCount() uint {
	return g.Calls
}

func (g *{{.Name}}) SetReturnValue(val {{.Type}}) {
	g.Returns = val
}

func (g *{{.Name}}) SetError(err error) {
	g.Error = err
}
`
