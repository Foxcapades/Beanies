// Generated at 2019-09-24T21:57:44-04:00
package bean

type IntSliceErrGetter struct {
	Calls   uint
	Returns []int
	Error		error
}

func (g *IntSliceErrGetter) Get() ([]int, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g IntSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *IntSliceErrGetter) SetReturnValue(val []int) {
	g.Returns = val
}

func (g *IntSliceErrGetter) SetError(err error) {
	g.Error = err
}
