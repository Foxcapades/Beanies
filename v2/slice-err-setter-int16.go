// Generated at 2019-09-24T21:57:48-04:00
package bean

type Int16SliceErrSetter struct {
	Calls   uint
	Inputs  [][]int16
	Error   error
}

func (g *Int16SliceErrSetter) Set(v []int16) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Int16SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int16SliceErrSetter) InputValues() [][]int16 {
	return g.Inputs
}

func (g *Int16SliceErrSetter) SetError(err error) {
	g.Error = err
}
