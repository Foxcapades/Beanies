// Generated at 2019-09-24T21:57:52-04:00
package bean

type RuneSliceErrGetter struct {
	Calls   uint
	Returns []rune
	Error		error
}

func (g *RuneSliceErrGetter) Get() ([]rune, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g RuneSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *RuneSliceErrGetter) SetReturnValue(val []rune) {
	g.Returns = val
}

func (g *RuneSliceErrGetter) SetError(err error) {
	g.Error = err
}
