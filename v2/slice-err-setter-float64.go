// Generated at 2019-09-24T21:49:00-04:00
package bean

type Float64SliceErrSetter struct {
	Calls   uint
	Inputs  [][]float64
}

func (g *Float64SliceErrSetter) Set(v []float64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Float64SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Float64SliceErrSetter) InputValues() [][]float64 {
	return g.Inputs
}
