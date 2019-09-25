// Generated at 2019-09-24T21:58:02-04:00
package bean

type Uint64SliceErrSetter struct {
	Calls   uint
	Inputs  [][]uint64
	Error   error
}

func (g *Uint64SliceErrSetter) Set(v []uint64) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Uint64SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint64SliceErrSetter) InputValues() [][]uint64 {
	return g.Inputs
}

func (g *Uint64SliceErrSetter) SetError(err error) {
	g.Error = err
}
