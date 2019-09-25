// Generated at 2019-09-24T21:48:57-04:00
package bean

type Complex128SliceSetter struct {
	Calls   uint
	Inputs  [][]complex128
}

func (g *Complex128SliceSetter) Set(v []complex128) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Complex128SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Complex128SliceSetter) InputValues() [][]complex128 {
	return g.Inputs
}
