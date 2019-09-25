// Generated at 2019-09-24T21:57:44-04:00
package bean

type IntSliceErrSetter struct {
	Calls   uint
	Inputs  [][]int
	Error   error
}

func (g *IntSliceErrSetter) Set(v []int) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g IntSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *IntSliceErrSetter) InputValues() [][]int {
	return g.Inputs
}

func (g *IntSliceErrSetter) SetError(err error) {
	g.Error = err
}
