// Generated at 2019-09-24T21:57:54-04:00
package bean

type StringSliceErrGetter struct {
	Calls   uint
	Returns []string
	Error		error
}

func (g *StringSliceErrGetter) Get() ([]string, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g StringSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *StringSliceErrGetter) SetReturnValue(val []string) {
	g.Returns = val
}

func (g *StringSliceErrGetter) SetError(err error) {
	g.Error = err
}
