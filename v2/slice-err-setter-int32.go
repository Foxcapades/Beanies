// Generated at 2019-09-24T21:49:06-04:00
package bean

type Int32SliceErrSetter struct {
	Calls   uint
	Inputs  [][]int32
}

func (g *Int32SliceErrSetter) Set(v []int32) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int32SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int32SliceErrSetter) InputValues() [][]int32 {
	return g.Inputs
}
