// Generated at 2019-09-24T21:57:42-04:00
package bean

type Float64SliceSetter struct {
	Calls   uint
	Inputs  [][]float64
}

func (g *Float64SliceSetter) Set(v []float64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Float64SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Float64SliceSetter) InputValues() [][]float64 {
	return g.Inputs
}
