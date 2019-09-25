// Generated at 2019-09-24T21:57:43-04:00
package bean

type Float64SliceErrSetter struct {
	Calls   uint
	Inputs  [][]float64
	Error   error
}

func (g *Float64SliceErrSetter) Set(v []float64) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Float64SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Float64SliceErrSetter) InputValues() [][]float64 {
	return g.Inputs
}

func (g *Float64SliceErrSetter) SetError(err error) {
	g.Error = err
}
