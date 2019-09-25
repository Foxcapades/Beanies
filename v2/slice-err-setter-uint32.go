// Generated at 2019-09-24T21:49:18-04:00
package bean

type Uint32SliceErrSetter struct {
	Calls   uint
	Inputs  [][]uint32
}

func (g *Uint32SliceErrSetter) Set(v []uint32) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint32SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint32SliceErrSetter) InputValues() [][]uint32 {
	return g.Inputs
}
