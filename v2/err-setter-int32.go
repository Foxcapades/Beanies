// Generated at 2019-09-24T21:57:48-04:00
package bean

type Int32ErrSetter struct {
	Calls   uint
	Inputs  []int32
	Error   error
}

func (g *Int32ErrSetter) Set(v int32) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Int32ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int32ErrSetter) InputValues() []int32 {
	return g.Inputs
}

func (g *Int32ErrSetter) SetError(err error) {
	g.Error = err
}
