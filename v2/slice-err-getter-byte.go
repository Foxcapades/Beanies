// Generated at 2019-09-24T21:48:53-04:00
package bean

type ByteSliceErrGetter struct {
	Calls   uint
	Returns []byte
}

func (g *ByteSliceErrGetter) Get() []byte {
	g.Calls++
	return g.Returns
}

func (g ByteSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *ByteSliceErrGetter) SetReturnValue(val []byte) {
	g.Returns = val
}
