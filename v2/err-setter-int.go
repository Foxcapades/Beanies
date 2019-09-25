// Generated at 2019-09-24T21:57:43-04:00
package bean

type IntErrSetter struct {
	Calls   uint
	Inputs  []int
	Error   error
}

func (g *IntErrSetter) Set(v int) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g IntErrSetter) CallCount() uint {
	return g.Calls
}

func (g *IntErrSetter) InputValues() []int {
	return g.Inputs
}

func (g *IntErrSetter) SetError(err error) {
	g.Error = err
}
