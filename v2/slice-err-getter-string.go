// Generated at 2019-09-24T21:49:11-04:00
package bean

type StringSliceErrGetter struct {
	Calls   uint
	Returns []string
}

func (g *StringSliceErrGetter) Get() []string {
	g.Calls++
	return g.Returns
}

func (g StringSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *StringSliceErrGetter) SetReturnValue(val []string) {
	g.Returns = val
}
