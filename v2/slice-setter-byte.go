// Generated at 2019-09-24T21:48:53-04:00
package bean

type ByteSliceSetter struct {
	Calls   uint
	Inputs  [][]byte
}

func (g *ByteSliceSetter) Set(v []byte) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g ByteSliceSetter) CallCount() uint {
	return g.Calls
}

func (g *ByteSliceSetter) InputValues() [][]byte {
	return g.Inputs
}
