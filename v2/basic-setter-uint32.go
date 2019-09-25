// Generated at 2019-09-24T21:49:17-04:00
package bean

type Uint32Setter struct {
	Calls   uint
	Inputs  []uint32
}

func (g *Uint32Setter) Set(v uint32) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint32Setter) CallCount() uint {
	return g.Calls
}

func (g *Uint32Setter) InputValues() []uint32 {
	return g.Inputs
}
