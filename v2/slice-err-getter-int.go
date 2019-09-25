// Generated at 2019-09-24T21:49:02-04:00
package bean

type IntSliceErrGetter struct {
	Calls   uint
	Returns []int
}

func (g *IntSliceErrGetter) Get() []int {
	g.Calls++
	return g.Returns
}

func (g IntSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *IntSliceErrGetter) SetReturnValue(val []int) {
	g.Returns = val
}
