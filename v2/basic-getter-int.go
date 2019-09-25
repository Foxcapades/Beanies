// Generated at 2019-09-24T21:57:43-04:00
package bean

type IntGetter struct {
	Calls   uint
	Returns int
}

func (g *IntGetter) Get() int {
	g.Calls++
	return g.Returns
}

func (g IntGetter) CallCount() uint {
	return g.Calls
}

func (g *IntGetter) SetReturnValue(val int) {
	g.Returns = val
}
