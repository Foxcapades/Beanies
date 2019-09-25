// Generated at 2019-09-24T21:49:13-04:00
package bean

type UintSliceErrSetter struct {
	Calls   uint
	Inputs  [][]uint
}

func (g *UintSliceErrSetter) Set(v []uint) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g UintSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *UintSliceErrSetter) InputValues() [][]uint {
	return g.Inputs
}
