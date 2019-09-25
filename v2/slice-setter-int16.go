// Generated at 2019-09-24T21:57:47-04:00
package bean

type Int16SliceSetter struct {
	Calls   uint
	Inputs  [][]int16
}

func (g *Int16SliceSetter) Set(v []int16) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int16SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Int16SliceSetter) InputValues() [][]int16 {
	return g.Inputs
}
