// Generated at 2019-09-24T21:48:57-04:00
package bean

type Complex128SliceErrGetter struct {
	Calls   uint
	Returns []complex128
}

func (g *Complex128SliceErrGetter) Get() []complex128 {
	g.Calls++
	return g.Returns
}

func (g Complex128SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Complex128SliceErrGetter) SetReturnValue(val []complex128) {
	g.Returns = val
}
