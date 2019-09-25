// Generated at 2019-09-24T21:49:03-04:00
package bean

type Int8ErrSetter struct {
	Calls   uint
	Inputs  []int8
	Error   error
}

func (g *Int8ErrSetter) Set(v int8) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Int8ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int8ErrSetter) InputValues() []int8 {
	return g.Inputs
}

func (g *Int8ErrSetter) SetError(err error) {
	g.Error = err
}
