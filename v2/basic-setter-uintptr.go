// Generated at 2019-09-24T21:49:20-04:00
package bean

type UintptrSetter struct {
	Calls   uint
	Inputs  []uintptr
}

func (g *UintptrSetter) Set(v uintptr) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g UintptrSetter) CallCount() uint {
	return g.Calls
}

func (g *UintptrSetter) InputValues() []uintptr {
	return g.Inputs
}
