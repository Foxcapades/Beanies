// Generated at 2019-09-24T21:57:58-04:00
package bean

type Uint16ErrSetter struct {
	Calls   uint
	Inputs  []uint16
	Error   error
}

func (g *Uint16ErrSetter) Set(v uint16) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Uint16ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint16ErrSetter) InputValues() []uint16 {
	return g.Inputs
}

func (g *Uint16ErrSetter) SetError(err error) {
	g.Error = err
}
