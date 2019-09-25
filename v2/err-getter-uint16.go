// Generated at 2019-09-24T21:49:15-04:00
package bean

type Uint16ErrGetter struct {
	Calls   uint
	Returns uint16
	Error		error
}

func (g *Uint16ErrGetter) Get() (uint16, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Uint16ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint16ErrGetter) SetReturnValue(val uint16) {
	g.Returns = val
}

func (g *Uint16ErrGetter) SetError(err error) {
	g.Error = err
}
