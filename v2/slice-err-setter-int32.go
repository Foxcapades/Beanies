// Generated at 2019-09-24T21:57:49-04:00
package bean

type Int32SliceErrSetter struct {
	Calls   uint
	Inputs  [][]int32
	Error   error
}

func (g *Int32SliceErrSetter) Set(v []int32) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Int32SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int32SliceErrSetter) InputValues() [][]int32 {
	return g.Inputs
}

func (g *Int32SliceErrSetter) SetError(err error) {
	g.Error = err
}
