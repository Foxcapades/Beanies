// Generated at 2019-09-24T21:48:55-04:00
package bean

type Complex64SliceErrGetter struct {
	Calls   uint
	Returns []complex64
}

func (g *Complex64SliceErrGetter) Get() []complex64 {
	g.Calls++
	return g.Returns
}

func (g Complex64SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Complex64SliceErrGetter) SetReturnValue(val []complex64) {
	g.Returns = val
}
