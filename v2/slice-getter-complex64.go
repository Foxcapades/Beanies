// Generated at 2019-09-24T21:57:37-04:00
package bean

type Complex64SliceGetter struct {
	Calls   uint
	Returns []complex64
}

func (g *Complex64SliceGetter) Get() []complex64 {
	g.Calls++
	return g.Returns
}

func (g Complex64SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Complex64SliceGetter) SetReturnValue(val []complex64) {
	g.Returns = val
}
