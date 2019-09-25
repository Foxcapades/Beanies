// Generated at 2019-09-24T21:48:59-04:00
package bean

type Float32SliceErrSetter struct {
	Calls   uint
	Inputs  [][]float32
}

func (g *Float32SliceErrSetter) Set(v []float32) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Float32SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Float32SliceErrSetter) InputValues() [][]float32 {
	return g.Inputs
}
