// Generated at 2019-09-24T21:48:54-04:00
package bean

type ByteSliceErrSetter struct {
	Calls   uint
	Inputs  [][]byte
}

func (g *ByteSliceErrSetter) Set(v []byte) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g ByteSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *ByteSliceErrSetter) InputValues() [][]byte {
	return g.Inputs
}
