// Generated at 2019-09-24T21:48:55-04:00
package bean

type Complex64SliceSetter struct {
	Calls   uint
	Inputs  [][]complex64
}

func (g *Complex64SliceSetter) Set(v []complex64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Complex64SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Complex64SliceSetter) InputValues() [][]complex64 {
	return g.Inputs
}
