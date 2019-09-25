// Generated at 2019-09-24T21:48:55-04:00
package bean

type Complex64SliceErrSetter struct {
	Calls   uint
	Inputs  [][]complex64
}

func (g *Complex64SliceErrSetter) Set(v []complex64) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Complex64SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Complex64SliceErrSetter) InputValues() [][]complex64 {
	return g.Inputs
}
