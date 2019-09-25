// Generated at 2019-09-24T21:49:10-04:00
package bean

type StringErrGetter struct {
	Calls   uint
	Returns string
	Error		error
}

func (g *StringErrGetter) Get() (string, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g StringErrGetter) CallCount() uint {
	return g.Calls
}

func (g *StringErrGetter) SetReturnValue(val string) {
	g.Returns = val
}

func (g *StringErrGetter) SetError(err error) {
	g.Error = err
}
