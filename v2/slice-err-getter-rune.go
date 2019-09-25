// Generated at 2019-09-24T21:49:09-04:00
package bean

type RuneSliceErrGetter struct {
	Calls   uint
	Returns []rune
}

func (g *RuneSliceErrGetter) Get() []rune {
	g.Calls++
	return g.Returns
}

func (g RuneSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *RuneSliceErrGetter) SetReturnValue(val []rune) {
	g.Returns = val
}
