// Generated at 2019-09-24T21:49:09-04:00
package bean

type RuneSliceSetter struct {
	Calls   uint
	Inputs  [][]rune
}

func (g *RuneSliceSetter) Set(v []rune) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g RuneSliceSetter) CallCount() uint {
	return g.Calls
}

func (g *RuneSliceSetter) InputValues() [][]rune {
	return g.Inputs
}
