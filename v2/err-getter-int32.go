// Generated at 2019-09-24T21:49:05-04:00
package bean

type Int32ErrGetter struct {
	Calls   uint
	Returns int32
	Error		error
}

func (g *Int32ErrGetter) Get() (int32, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Int32ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int32ErrGetter) SetReturnValue(val int32) {
	g.Returns = val
}

func (g *Int32ErrGetter) SetError(err error) {
	g.Error = err
}
