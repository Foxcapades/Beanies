// Generated at 2019-09-24T21:49:12-04:00
package bean

type UintErrGetter struct {
	Calls   uint
	Returns uint
	Error		error
}

func (g *UintErrGetter) Get() (uint, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g UintErrGetter) CallCount() uint {
	return g.Calls
}

func (g *UintErrGetter) SetReturnValue(val uint) {
	g.Returns = val
}

func (g *UintErrGetter) SetError(err error) {
	g.Error = err
}
