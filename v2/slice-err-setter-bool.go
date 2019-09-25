// Generated at 2019-09-24T21:57:35-04:00
package bean

type BoolSliceErrSetter struct {
	Calls   uint
	Inputs  [][]bool
	Error   error
}

func (g *BoolSliceErrSetter) Set(v []bool) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g BoolSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *BoolSliceErrSetter) InputValues() [][]bool {
	return g.Inputs
}

func (g *BoolSliceErrSetter) SetError(err error) {
	g.Error = err
}
