// Generated at 2019-09-24T21:57:52-04:00
package bean

type RuneSliceGetter struct {
	Calls   uint
	Returns []rune
}

func (g *RuneSliceGetter) Get() []rune {
	g.Calls++
	return g.Returns
}

func (g RuneSliceGetter) CallCount() uint {
	return g.Calls
}

func (g *RuneSliceGetter) SetReturnValue(val []rune) {
	g.Returns = val
}
