// Generated at 2019-09-24T21:57:39-04:00
package bean

type Complex128ErrSetter struct {
	Calls   uint
	Inputs  []complex128
	Error   error
}

func (g *Complex128ErrSetter) Set(v complex128) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Complex128ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Complex128ErrSetter) InputValues() []complex128 {
	return g.Inputs
}

func (g *Complex128ErrSetter) SetError(err error) {
	g.Error = err
}
