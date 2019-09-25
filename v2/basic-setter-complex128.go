// Generated at 2019-09-24T21:48:56-04:00
package bean

type Complex128Setter struct {
	Calls   uint
	Inputs  []complex128
}

func (g *Complex128Setter) Set(v complex128) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Complex128Setter) CallCount() uint {
	return g.Calls
}

func (g *Complex128Setter) InputValues() []complex128 {
	return g.Inputs
}
