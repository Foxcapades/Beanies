// Generated at 2019-09-24T21:48:58-04:00
package bean

type Float32ErrSetter struct {
	Calls   uint
	Inputs  []float32
	Error   error
}

func (g *Float32ErrSetter) Set(v float32) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Float32ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Float32ErrSetter) InputValues() []float32 {
	return g.Inputs
}

func (g *Float32ErrSetter) SetError(err error) {
	g.Error = err
}
