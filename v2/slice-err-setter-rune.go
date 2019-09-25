// Generated at 2019-09-24T21:49:10-04:00
package bean

type RuneSliceErrSetter struct {
	Calls   uint
	Inputs  [][]rune
}

func (g *RuneSliceErrSetter) Set(v []rune) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g RuneSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *RuneSliceErrSetter) InputValues() [][]rune {
	return g.Inputs
}
