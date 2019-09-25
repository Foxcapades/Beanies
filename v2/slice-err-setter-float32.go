// Generated at 2019-09-24T21:57:41-04:00
package bean

type Float32SliceErrSetter struct {
	Calls   uint
	Inputs  [][]float32
	Error   error
}

func (g *Float32SliceErrSetter) Set(v []float32) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Float32SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Float32SliceErrSetter) InputValues() [][]float32 {
	return g.Inputs
}

func (g *Float32SliceErrSetter) SetError(err error) {
	g.Error = err
}
