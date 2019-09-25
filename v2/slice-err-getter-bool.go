// Generated at 2019-09-24T21:48:52-04:00
package bean

type BoolSliceErrGetter struct {
	Calls   uint
	Returns []bool
}

func (g *BoolSliceErrGetter) Get() []bool {
	g.Calls++
	return g.Returns
}

func (g BoolSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *BoolSliceErrGetter) SetReturnValue(val []bool) {
	g.Returns = val
}
