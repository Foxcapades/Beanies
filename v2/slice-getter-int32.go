// Generated at 2019-09-24T21:57:49-04:00
package bean

type Int32SliceGetter struct {
	Calls   uint
	Returns []int32
}

func (g *Int32SliceGetter) Get() []int32 {
	g.Calls++
	return g.Returns
}

func (g Int32SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Int32SliceGetter) SetReturnValue(val []int32) {
	g.Returns = val
}
