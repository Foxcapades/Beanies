// Generated at 2019-09-24T21:49:18-04:00
package bean

type Uint64Setter struct {
	Calls   uint
	Inputs  []uint64
}

func (g *Uint64Setter) Set(v uint64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint64Setter) CallCount() uint {
	return g.Calls
}

func (g *Uint64Setter) InputValues() []uint64 {
	return g.Inputs
}
