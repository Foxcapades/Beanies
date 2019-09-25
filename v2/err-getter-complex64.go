// Generated at 2019-09-24T21:48:54-04:00
package bean

type Complex64ErrGetter struct {
	Calls   uint
	Returns complex64
	Error		error
}

func (g *Complex64ErrGetter) Get() (complex64, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Complex64ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Complex64ErrGetter) SetReturnValue(val complex64) {
	g.Returns = val
}

func (g *Complex64ErrGetter) SetError(err error) {
	g.Error = err
}
