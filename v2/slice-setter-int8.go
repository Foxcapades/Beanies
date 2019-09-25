// Generated at 2019-09-24T21:57:46-04:00
package bean

type Int8SliceSetter struct {
	Calls   uint
	Inputs  [][]int8
}

func (g *Int8SliceSetter) Set(v []int8) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int8SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Int8SliceSetter) InputValues() [][]int8 {
	return g.Inputs
}
