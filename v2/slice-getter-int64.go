// Generated at 2019-09-24T21:57:50-04:00
package bean

type Int64SliceGetter struct {
	Calls   uint
	Returns []int64
}

func (g *Int64SliceGetter) Get() []int64 {
	g.Calls++
	return g.Returns
}

func (g Int64SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Int64SliceGetter) SetReturnValue(val []int64) {
	g.Returns = val
}
