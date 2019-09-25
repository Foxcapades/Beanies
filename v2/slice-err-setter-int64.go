// Generated at 2019-09-24T21:49:08-04:00
package bean

type Int64SliceErrSetter struct {
	Calls   uint
	Inputs  [][]int64
}

func (g *Int64SliceErrSetter) Set(v []int64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int64SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int64SliceErrSetter) InputValues() [][]int64 {
	return g.Inputs
}
