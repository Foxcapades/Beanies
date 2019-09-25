// Generated at 2019-09-24T21:58:04-04:00
package bean

type UintptrSliceErrSetter struct {
	Calls   uint
	Inputs  [][]uintptr
	Error   error
}

func (g *UintptrSliceErrSetter) Set(v []uintptr) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g UintptrSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *UintptrSliceErrSetter) InputValues() [][]uintptr {
	return g.Inputs
}

func (g *UintptrSliceErrSetter) SetError(err error) {
	g.Error = err
}
