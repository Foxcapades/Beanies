// Generated at 2019-09-24T21:57:57-04:00
package bean

type Uint8SliceErrSetter struct {
	Calls   uint
	Inputs  [][]uint8
	Error   error
}

func (g *Uint8SliceErrSetter) Set(v []uint8) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Uint8SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint8SliceErrSetter) InputValues() [][]uint8 {
	return g.Inputs
}

func (g *Uint8SliceErrSetter) SetError(err error) {
	g.Error = err
}
