// Generated at 2019-09-24T21:49:06-04:00
package bean

type Int32SliceSetter struct {
	Calls   uint
	Inputs  [][]int32
}

func (g *Int32SliceSetter) Set(v []int32) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int32SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Int32SliceSetter) InputValues() [][]int32 {
	return g.Inputs
}
