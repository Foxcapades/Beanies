// Generated at 2019-09-24T21:49:21-04:00
package bean

type UintptrSliceErrSetter struct {
	Calls   uint
	Inputs  [][]uintptr
}

func (g *UintptrSliceErrSetter) Set(v []uintptr) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g UintptrSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *UintptrSliceErrSetter) InputValues() [][]uintptr {
	return g.Inputs
}
