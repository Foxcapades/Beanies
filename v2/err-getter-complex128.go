// Generated at 2019-09-24T21:48:56-04:00
package bean

type Complex128ErrGetter struct {
	Calls   uint
	Returns complex128
	Error		error
}

func (g *Complex128ErrGetter) Get() (complex128, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Complex128ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Complex128ErrGetter) SetReturnValue(val complex128) {
	g.Returns = val
}

func (g *Complex128ErrGetter) SetError(err error) {
	g.Error = err
}
