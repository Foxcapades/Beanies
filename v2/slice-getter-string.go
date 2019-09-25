// Generated at 2019-09-24T21:49:11-04:00
package bean

type StringSliceGetter struct {
	Calls   uint
	Returns []string
}

func (g *StringSliceGetter) Get() []string {
	g.Calls++
	return g.Returns
}

func (g StringSliceGetter) CallCount() uint {
	return g.Calls
}

func (g *StringSliceGetter) SetReturnValue(val []string) {
	g.Returns = val
}
