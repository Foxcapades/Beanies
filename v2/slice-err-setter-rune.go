// Generated at 2019-09-24T21:57:52-04:00
package bean

type RuneSliceErrSetter struct {
	Calls   uint
	Inputs  [][]rune
	Error   error
}

func (g *RuneSliceErrSetter) Set(v []rune) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g RuneSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *RuneSliceErrSetter) InputValues() [][]rune {
	return g.Inputs
}

func (g *RuneSliceErrSetter) SetError(err error) {
	g.Error = err
}
