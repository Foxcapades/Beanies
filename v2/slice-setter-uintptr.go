// Generated at 2019-09-24T21:49:21-04:00
package bean

type UintptrSliceSetter struct {
	Calls   uint
	Inputs  [][]uintptr
}

func (g *UintptrSliceSetter) Set(v []uintptr) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g UintptrSliceSetter) CallCount() uint {
	return g.Calls
}

func (g *UintptrSliceSetter) InputValues() [][]uintptr {
	return g.Inputs
}
