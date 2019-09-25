// Generated at 2019-09-24T21:49:05-04:00
package bean

type Int16SliceErrSetter struct {
	Calls   uint
	Inputs  [][]int16
}

func (g *Int16SliceErrSetter) Set(v []int16) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int16SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int16SliceErrSetter) InputValues() [][]int16 {
	return g.Inputs
}
