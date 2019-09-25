// Generated at 2019-09-24T21:57:56-04:00
package bean

type UintSliceErrSetter struct {
	Calls   uint
	Inputs  [][]uint
	Error   error
}

func (g *UintSliceErrSetter) Set(v []uint) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g UintSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *UintSliceErrSetter) InputValues() [][]uint {
	return g.Inputs
}

func (g *UintSliceErrSetter) SetError(err error) {
	g.Error = err
}
