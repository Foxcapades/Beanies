// Generated at 2019-09-24T21:48:54-04:00
package bean

type Complex64Setter struct {
	Calls   uint
	Inputs  []complex64
}

func (g *Complex64Setter) Set(v complex64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Complex64Setter) CallCount() uint {
	return g.Calls
}

func (g *Complex64Setter) InputValues() []complex64 {
	return g.Inputs
}
