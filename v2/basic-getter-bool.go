// Generated at 2019-09-24T21:48:50-04:00
package bean

type BoolGetter struct {
	Calls   uint
	Returns bool
}

func (g *BoolGetter) Get() bool {
	g.Calls++
	return g.Returns
}

func (g BoolGetter) CallCount() uint {
	return g.Calls
}

func (g *BoolGetter) SetReturnValue(val bool) {
	g.Returns = val
}
