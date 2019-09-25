// Generated at 2019-09-24T21:57:46-04:00
package bean

type Int8SliceErrSetter struct {
	Calls   uint
	Inputs  [][]int8
	Error   error
}

func (g *Int8SliceErrSetter) Set(v []int8) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Int8SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int8SliceErrSetter) InputValues() [][]int8 {
	return g.Inputs
}

func (g *Int8SliceErrSetter) SetError(err error) {
	g.Error = err
}
