// Generated at 2019-09-24T21:57:36-04:00
package bean

type ByteSliceErrSetter struct {
	Calls   uint
	Inputs  [][]byte
	Error   error
}

func (g *ByteSliceErrSetter) Set(v []byte) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g ByteSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *ByteSliceErrSetter) InputValues() [][]byte {
	return g.Inputs
}

func (g *ByteSliceErrSetter) SetError(err error) {
	g.Error = err
}
