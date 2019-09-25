// Generated at 2019-09-24T21:49:03-04:00
package bean

type Int8SliceErrSetter struct {
	Calls   uint
	Inputs  [][]int8
}

func (g *Int8SliceErrSetter) Set(v []int8) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int8SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int8SliceErrSetter) InputValues() [][]int8 {
	return g.Inputs
}