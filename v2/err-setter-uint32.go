// Generated at 2019-09-24T21:58:00-04:00
package bean

type Uint32ErrSetter struct {
	Calls   uint
	Inputs  []uint32
	Error   error
}

func (g *Uint32ErrSetter) Set(v uint32) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Uint32ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint32ErrSetter) InputValues() []uint32 {
	return g.Inputs
}

func (g *Uint32ErrSetter) SetError(err error) {
	g.Error = err
}
