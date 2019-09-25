// Generated at 2019-09-24T21:57:36-04:00
package bean

type ByteSliceErrGetter struct {
	Calls   uint
	Returns []byte
	Error		error
}

func (g *ByteSliceErrGetter) Get() ([]byte, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g ByteSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *ByteSliceErrGetter) SetReturnValue(val []byte) {
	g.Returns = val
}

func (g *ByteSliceErrGetter) SetError(err error) {
	g.Error = err
}
