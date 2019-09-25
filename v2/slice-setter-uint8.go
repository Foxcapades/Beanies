// Generated at 2019-09-24T21:57:57-04:00
package bean

type Uint8SliceSetter struct {
	Calls   uint
	Inputs  [][]uint8
}

func (g *Uint8SliceSetter) Set(v []uint8) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint8SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint8SliceSetter) InputValues() [][]uint8 {
	return g.Inputs
}
