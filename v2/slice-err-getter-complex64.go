// Generated at 2019-09-24T21:57:38-04:00
package bean

type Complex64SliceErrGetter struct {
	Calls   uint
	Returns []complex64
	Error		error
}

func (g *Complex64SliceErrGetter) Get() ([]complex64, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Complex64SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Complex64SliceErrGetter) SetReturnValue(val []complex64) {
	g.Returns = val
}

func (g *Complex64SliceErrGetter) SetError(err error) {
	g.Error = err
}
