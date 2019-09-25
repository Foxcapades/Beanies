// Generated at 2019-09-24T21:58:01-04:00
package bean

type Uint64ErrGetter struct {
	Calls   uint
	Returns uint64
	Error		error
}

func (g *Uint64ErrGetter) Get() (uint64, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Uint64ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint64ErrGetter) SetReturnValue(val uint64) {
	g.Returns = val
}

func (g *Uint64ErrGetter) SetError(err error) {
	g.Error = err
}
