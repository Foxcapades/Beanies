// Generated at 2019-09-24T21:57:41-04:00
package bean

type Float32SliceSetter struct {
	Calls   uint
	Inputs  [][]float32
}

func (g *Float32SliceSetter) Set(v []float32) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Float32SliceSetter) CallCount() uint {
	return g.Calls
}

func (g *Float32SliceSetter) InputValues() [][]float32 {
	return g.Inputs
}
