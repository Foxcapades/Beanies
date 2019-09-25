// Generated at 2019-09-24T21:57:44-04:00
package bean

type IntSliceGetter struct {
	Calls   uint
	Returns []int
}

func (g *IntSliceGetter) Get() []int {
	g.Calls++
	return g.Returns
}

func (g IntSliceGetter) CallCount() uint {
	return g.Calls
}

func (g *IntSliceGetter) SetReturnValue(val []int) {
	g.Returns = val
}
