// Generated at 2019-09-24T21:57:35-04:00
package bean

type ByteErrSetter struct {
	Calls   uint
	Inputs  []byte
	Error   error
}

func (g *ByteErrSetter) Set(v byte) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g ByteErrSetter) CallCount() uint {
	return g.Calls
}

func (g *ByteErrSetter) InputValues() []byte {
	return g.Inputs
}

func (g *ByteErrSetter) SetError(err error) {
	g.Error = err
}
