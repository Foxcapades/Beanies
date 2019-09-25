// Generated at 2019-09-24T21:48:51-04:00
package bean

type BoolSliceSetter struct {
	Calls   uint
	Inputs  [][]bool
}

func (g *BoolSliceSetter) Set(v []bool) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g BoolSliceSetter) CallCount() uint {
	return g.Calls
}

func (g *BoolSliceSetter) InputValues() [][]bool {
	return g.Inputs
}
