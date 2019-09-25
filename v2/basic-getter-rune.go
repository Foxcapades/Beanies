// Generated at 2019-09-24T21:57:51-04:00
package bean

type RuneGetter struct {
	Calls   uint
	Returns rune
}

func (g *RuneGetter) Get() rune {
	g.Calls++
	return g.Returns
}

func (g RuneGetter) CallCount() uint {
	return g.Calls
}

func (g *RuneGetter) SetReturnValue(val rune) {
	g.Returns = val
}
