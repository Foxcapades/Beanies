// Generated at 2019-09-24T21:48:53-04:00
package bean

type ByteSliceGetter struct {
	Calls   uint
	Returns []byte
}

func (g *ByteSliceGetter) Get() []byte {
	g.Calls++
	return g.Returns
}

func (g ByteSliceGetter) CallCount() uint {
	return g.Calls
}

func (g *ByteSliceGetter) SetReturnValue(val []byte) {
	g.Returns = val
}
