// Generated at 2019-09-24T21:49:20-04:00
package bean

type Uint64SliceErrSetter struct {
	Calls   uint
	Inputs  [][]uint64
}

func (g *Uint64SliceErrSetter) Set(v []uint64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint64SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint64SliceErrSetter) InputValues() [][]uint64 {
	return g.Inputs
}
