// Generated at 2019-09-24T21:49:12-04:00
package bean

type UintErrSetter struct {
	Calls   uint
	Inputs  []uint
	Error   error
}

func (g *UintErrSetter) Set(v uint) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g UintErrSetter) CallCount() uint {
	return g.Calls
}

func (g *UintErrSetter) InputValues() []uint {
	return g.Inputs
}

func (g *UintErrSetter) SetError(err error) {
	g.Error = err
}
