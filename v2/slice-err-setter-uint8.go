// Generated at 2019-09-24T21:49:15-04:00
package bean

type Uint8SliceErrSetter struct {
	Calls   uint
	Inputs  [][]uint8
}

func (g *Uint8SliceErrSetter) Set(v []uint8) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint8SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint8SliceErrSetter) InputValues() [][]uint8 {
	return g.Inputs
}
