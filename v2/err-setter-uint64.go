// Generated at 2019-09-24T21:49:19-04:00
package bean

type Uint64ErrSetter struct {
	Calls   uint
	Inputs  []uint64
	Error   error
}

func (g *Uint64ErrSetter) Set(v uint64) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Uint64ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint64ErrSetter) InputValues() []uint64 {
	return g.Inputs
}

func (g *Uint64ErrSetter) SetError(err error) {
	g.Error = err
}
