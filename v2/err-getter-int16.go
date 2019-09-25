// Generated at 2019-09-24T21:49:04-04:00
package bean

type Int16ErrGetter struct {
	Calls   uint
	Returns int16
	Error		error
}

func (g *Int16ErrGetter) Get() (int16, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Int16ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int16ErrGetter) SetReturnValue(val int16) {
	g.Returns = val
}

func (g *Int16ErrGetter) SetError(err error) {
	g.Error = err
}
