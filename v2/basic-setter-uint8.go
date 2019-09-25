// Generated at 2019-09-24T21:57:56-04:00
package bean

type Uint8Setter struct {
	Calls   uint
	Inputs  []uint8
}

func (g *Uint8Setter) Set(v uint8) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint8Setter) CallCount() uint {
	return g.Calls
}

func (g *Uint8Setter) InputValues() []uint8 {
	return g.Inputs
}
