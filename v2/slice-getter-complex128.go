// Generated at 2019-09-24T21:57:39-04:00
package bean

type Complex128SliceGetter struct {
	Calls   uint
	Returns []complex128
}

func (g *Complex128SliceGetter) Get() []complex128 {
	g.Calls++
	return g.Returns
}

func (g Complex128SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Complex128SliceGetter) SetReturnValue(val []complex128) {
	g.Returns = val
}
