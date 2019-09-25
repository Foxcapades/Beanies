// Generated at 2019-09-24T21:49:19-04:00
package bean

type Uint64SliceSetter struct {
	Calls   uint
	Inputs  [][]uint64
}

func (g *Uint64SliceSetter) Set(v []uint64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint64SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint64SliceSetter) InputValues() [][]uint64 {
	return g.Inputs
}
