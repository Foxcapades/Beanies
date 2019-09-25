// Generated at 2019-09-24T21:57:52-04:00
package bean

type RuneErrSetter struct {
	Calls   uint
	Inputs  []rune
	Error   error
}

func (g *RuneErrSetter) Set(v rune) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g RuneErrSetter) CallCount() uint {
	return g.Calls
}

func (g *RuneErrSetter) InputValues() []rune {
	return g.Inputs
}

func (g *RuneErrSetter) SetError(err error) {
	g.Error = err
}
