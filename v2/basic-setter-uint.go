// Generated at 2019-09-24T21:57:54-04:00
package bean

type UintSetter struct {
	Calls   uint
	Inputs  []uint
}

func (g *UintSetter) Set(v uint) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g UintSetter) CallCount() uint {
	return g.Calls
}

func (g *UintSetter) InputValues() []uint {
	return g.Inputs
}
