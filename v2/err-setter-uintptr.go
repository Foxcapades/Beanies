// Generated at 2019-09-24T21:49:20-04:00
package bean

type UintptrErrSetter struct {
	Calls   uint
	Inputs  []uintptr
	Error   error
}

func (g *UintptrErrSetter) Set(v uintptr) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g UintptrErrSetter) CallCount() uint {
	return g.Calls
}

func (g *UintptrErrSetter) InputValues() []uintptr {
	return g.Inputs
}

func (g *UintptrErrSetter) SetError(err error) {
	g.Error = err
}
