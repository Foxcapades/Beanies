// Generated at 2019-09-24T21:57:55-04:00
package bean

type UintSliceSetter struct {
	Calls   uint
	Inputs  [][]uint
}

func (g *UintSliceSetter) Set(v []uint) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g UintSliceSetter) CallCount() uint {
	return g.Calls
}

func (g *UintSliceSetter) InputValues() [][]uint {
	return g.Inputs
}
