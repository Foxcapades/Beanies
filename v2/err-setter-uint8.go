// Generated at 2019-09-24T21:49:14-04:00
package bean

type Uint8ErrSetter struct {
	Calls   uint
	Inputs  []uint8
	Error   error
}

func (g *Uint8ErrSetter) Set(v uint8) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g Uint8ErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint8ErrSetter) InputValues() []uint8 {
	return g.Inputs
}

func (g *Uint8ErrSetter) SetError(err error) {
	g.Error = err
}
