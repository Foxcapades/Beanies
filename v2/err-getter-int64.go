// Generated at 2019-09-24T21:57:50-04:00
package bean

type Int64ErrGetter struct {
	Calls   uint
	Returns int64
	Error		error
}

func (g *Int64ErrGetter) Get() (int64, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Int64ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int64ErrGetter) SetReturnValue(val int64) {
	g.Returns = val
}

func (g *Int64ErrGetter) SetError(err error) {
	g.Error = err
}
