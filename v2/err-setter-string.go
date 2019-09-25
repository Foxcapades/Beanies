// Generated at 2019-09-24T21:49:10-04:00
package bean

type StringErrSetter struct {
	Calls   uint
	Inputs  []string
	Error   error
}

func (g *StringErrSetter) Set(v string) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g StringErrSetter) CallCount() uint {
	return g.Calls
}

func (g *StringErrSetter) InputValues() []string {
	return g.Inputs
}

func (g *StringErrSetter) SetError(err error) {
	g.Error = err
}
