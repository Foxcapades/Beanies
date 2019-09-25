// Generated at 2019-09-24T21:57:58-04:00
package bean

type Uint16Setter struct {
	Calls   uint
	Inputs  []uint16
}

func (g *Uint16Setter) Set(v uint16) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint16Setter) CallCount() uint {
	return g.Calls
}

func (g *Uint16Setter) InputValues() []uint16 {
	return g.Inputs
}
