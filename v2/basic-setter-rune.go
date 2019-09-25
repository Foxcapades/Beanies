// Generated at 2019-09-24T21:57:51-04:00
package bean

type RuneSetter struct {
	Calls   uint
	Inputs  []rune
}

func (g *RuneSetter) Set(v rune) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g RuneSetter) CallCount() uint {
	return g.Calls
}

func (g *RuneSetter) InputValues() []rune {
	return g.Inputs
}
