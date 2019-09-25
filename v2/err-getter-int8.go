// Generated at 2019-09-24T21:49:02-04:00
package bean

type Int8ErrGetter struct {
	Calls   uint
	Returns int8
	Error		error
}

func (g *Int8ErrGetter) Get() (int8, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Int8ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int8ErrGetter) SetReturnValue(val int8) {
	g.Returns = val
}

func (g *Int8ErrGetter) SetError(err error) {
	g.Error = err
}
