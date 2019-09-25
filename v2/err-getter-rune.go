// Generated at 2019-09-24T21:57:51-04:00
package bean

type RuneErrGetter struct {
	Calls   uint
	Returns rune
	Error		error
}

func (g *RuneErrGetter) Get() (rune, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g RuneErrGetter) CallCount() uint {
	return g.Calls
}

func (g *RuneErrGetter) SetReturnValue(val rune) {
	g.Returns = val
}

func (g *RuneErrGetter) SetError(err error) {
	g.Error = err
}
