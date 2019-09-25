// Generated at 2019-09-24T21:57:43-04:00
package bean

type IntErrGetter struct {
	Calls   uint
	Returns int
	Error		error
}

func (g *IntErrGetter) Get() (int, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g IntErrGetter) CallCount() uint {
	return g.Calls
}

func (g *IntErrGetter) SetReturnValue(val int) {
	g.Returns = val
}

func (g *IntErrGetter) SetError(err error) {
	g.Error = err
}
