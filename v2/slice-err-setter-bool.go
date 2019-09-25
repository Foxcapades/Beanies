// Generated at 2019-09-24T21:48:52-04:00
package bean

type BoolSliceErrSetter struct {
	Calls   uint
	Inputs  [][]bool
}

func (g *BoolSliceErrSetter) Set(v []bool) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g BoolSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *BoolSliceErrSetter) InputValues() [][]bool {
	return g.Inputs
}
