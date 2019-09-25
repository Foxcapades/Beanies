// Generated at 2019-09-24T21:57:50-04:00
package bean

type Int64ErrSetter struct {
	Calls   uint
	Inputs  []int64
	Error   error
}

func (g *Int64ErrSetter) Set(v int64) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Int64ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Int64ErrSetter) InputValues() []int64 {
	return g.Inputs
}

func (g *Int64ErrSetter) SetError(err error) {
	g.Error = err
}
