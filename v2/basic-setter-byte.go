// Generated at 2019-09-24T21:48:52-04:00
package bean

type ByteSetter struct {
	Calls   uint
	Inputs  []byte
}

func (g *ByteSetter) Set(v byte) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g ByteSetter) CallCount() uint {
	return g.Calls
}

func (g *ByteSetter) InputValues() []byte {
	return g.Inputs
}
