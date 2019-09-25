// Generated at 2019-09-24T21:48:51-04:00
package bean

type BoolSliceGetter struct {
	Calls   uint
	Returns []bool
}

func (g *BoolSliceGetter) Get() []bool {
	g.Calls++
	return g.Returns
}

func (g BoolSliceGetter) CallCount() uint {
	return g.Calls
}

func (g *BoolSliceGetter) SetReturnValue(val []bool) {
	g.Returns = val
}
