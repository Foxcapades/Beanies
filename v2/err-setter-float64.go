// Generated at 2019-09-24T21:48:59-04:00
package bean

type Float64ErrSetter struct {
	Calls   uint
	Inputs  []float64
	Error   error
}

func (g *Float64ErrSetter) Set(v float64) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Float64ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Float64ErrSetter) InputValues() []float64 {
	return g.Inputs
}

func (g *Float64ErrSetter) SetError(err error) {
	g.Error = err
}
