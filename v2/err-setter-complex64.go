// Generated at 2019-09-24T21:57:37-04:00
package bean

type Complex64ErrSetter struct {
	Calls   uint
	Inputs  []complex64
	Error   error
}

func (g *Complex64ErrSetter) Set(v complex64) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Complex64ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Complex64ErrSetter) InputValues() []complex64 {
	return g.Inputs
}

func (g *Complex64ErrSetter) SetError(err error) {
	g.Error = err
}
