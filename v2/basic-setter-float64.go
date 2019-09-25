// Generated at 2019-09-24T21:57:41-04:00
package bean

type Float64Setter struct {
	Calls   uint
	Inputs  []float64
}

func (g *Float64Setter) Set(v float64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Float64Setter) CallCount() uint {
	return g.Calls
}

func (g *Float64Setter) InputValues() []float64 {
	return g.Inputs
}
