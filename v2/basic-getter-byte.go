// Generated at 2019-09-24T21:48:52-04:00
package bean

type ByteGetter struct {
	Calls   uint
	Returns byte
}

func (g *ByteGetter) Get() byte {
	g.Calls++
	return g.Returns
}

func (g ByteGetter) CallCount() uint {
	return g.Calls
}

func (g *ByteGetter) SetReturnValue(val byte) {
	g.Returns = val
}
