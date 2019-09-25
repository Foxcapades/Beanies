// Generated at 2019-09-24T21:57:40-04:00
package bean

type Float32Setter struct {
	Calls   uint
	Inputs  []float32
}

func (g *Float32Setter) Set(v float32) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Float32Setter) CallCount() uint {
	return g.Calls
}

func (g *Float32Setter) InputValues() []float32 {
	return g.Inputs
}
