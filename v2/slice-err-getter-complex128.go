// Generated at 2019-09-24T21:57:39-04:00
package bean

type Complex128SliceErrGetter struct {
	Calls   uint
	Returns []complex128
	Error		error
}

func (g *Complex128SliceErrGetter) Get() ([]complex128, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Complex128SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Complex128SliceErrGetter) SetReturnValue(val []complex128) {
	g.Returns = val
}

func (g *Complex128SliceErrGetter) SetError(err error) {
	g.Error = err
}
