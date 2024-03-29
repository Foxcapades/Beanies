// Generated at 2019-09-24T21:57:50-04:00
package bean

type Int64SliceSetter struct {
	Calls   uint
	Inputs  [][]int64
}

func (g *Int64SliceSetter) Set(v []int64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int64SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Int64SliceSetter) InputValues() [][]int64 {
	return g.Inputs
}
