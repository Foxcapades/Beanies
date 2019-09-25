// Generated at 2019-09-24T21:48:51-04:00
package bean

type BoolErrGetter struct {
	Calls   uint
	Returns bool
	Error		error
}

func (g *BoolErrGetter) Get() (bool, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g BoolErrGetter) CallCount() uint {
	return g.Calls
}

func (g *BoolErrGetter) SetReturnValue(val bool) {
	g.Returns = val
}

func (g *BoolErrGetter) SetError(err error) {
	g.Error = err
}
