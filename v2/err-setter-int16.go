// Generated at 2019-09-24T21:49:04-04:00
package bean

type Int16ErrSetter struct {
	Calls   uint
	Inputs  []int16
	Error   error
}

func (g *Int16ErrSetter) Set(v int16) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Int16ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int16ErrSetter) InputValues() []int16 {
	return g.Inputs
}

func (g *Int16ErrSetter) SetError(err error) {
	g.Error = err
}
