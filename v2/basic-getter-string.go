// Generated at 2019-09-24T21:57:53-04:00
package bean

type StringGetter struct {
	Calls   uint
	Returns string
}

func (g *StringGetter) Get() string {
	g.Calls++
	return g.Returns
}

func (g StringGetter) CallCount() uint {
	return g.Calls
}

func (g *StringGetter) SetReturnValue(val string) {
	g.Returns = val
}
