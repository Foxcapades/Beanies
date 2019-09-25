// Generated at 2019-09-24T21:48:51-04:00
package bean

type BoolErrSetter struct {
	Calls   uint
	Inputs  []bool
	Error   error
}

func (g *BoolErrSetter) Set(v bool) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g BoolErrSetter) CallCount() uint {
	return g.Calls
}

func (g *BoolErrSetter) InputValues() []bool {
	return g.Inputs
}

func (g *BoolErrSetter) SetError(err error) {
	g.Error = err
}
