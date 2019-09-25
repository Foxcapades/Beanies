// Generated at 2019-09-24T21:57:59-04:00
package bean

type Uint32ErrGetter struct {
	Calls   uint
	Returns uint32
	Error		error
}

func (g *Uint32ErrGetter) Get() (uint32, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Uint32ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint32ErrGetter) SetReturnValue(val uint32) {
	g.Returns = val
}

func (g *Uint32ErrGetter) SetError(err error) {
	g.Error = err
}
