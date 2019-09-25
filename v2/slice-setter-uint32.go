// Generated at 2019-09-24T21:49:18-04:00
package bean

type Uint32SliceSetter struct {
	Calls   uint
	Inputs  [][]uint32
}

func (g *Uint32SliceSetter) Set(v []uint32) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint32SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint32SliceSetter) InputValues() [][]uint32 {
	return g.Inputs
}
