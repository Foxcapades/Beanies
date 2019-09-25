// Generated at 2019-09-24T21:57:38-04:00
package bean

type Complex64SliceErrSetter struct {
	Calls   uint
	Inputs  [][]complex64
	Error   error
}

func (g *Complex64SliceErrSetter) Set(v []complex64) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Complex64SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Complex64SliceErrSetter) InputValues() [][]complex64 {
	return g.Inputs
}

func (g *Complex64SliceErrSetter) SetError(err error) {
	g.Error = err
}
