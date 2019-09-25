// Generated at 2019-09-24T21:57:56-04:00
package bean

type Uint8ErrGetter struct {
	Calls   uint
	Returns uint8
	Error		error
}

func (g *Uint8ErrGetter) Get() (uint8, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Uint8ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint8ErrGetter) SetReturnValue(val uint8) {
	g.Returns = val
}

func (g *Uint8ErrGetter) SetError(err error) {
	g.Error = err
}
