// Generated at 2019-09-24T21:48:53-04:00
package bean

type ByteErrGetter struct {
	Calls   uint
	Returns byte
	Error		error
}

func (g *ByteErrGetter) Get() (byte, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g ByteErrGetter) CallCount() uint {
	return g.Calls
}

func (g *ByteErrGetter) SetReturnValue(val byte) {
	g.Returns = val
}

func (g *ByteErrGetter) SetError(err error) {
	g.Error = err
}
