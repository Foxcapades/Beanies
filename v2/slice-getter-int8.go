// Generated at 2019-09-24T21:49:03-04:00
package bean

type Int8SliceGetter struct {
	Calls   uint
	Returns []int8
}

func (g *Int8SliceGetter) Get() []int8 {
	g.Calls++
	return g.Returns
}

func (g Int8SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Int8SliceGetter) SetReturnValue(val []int8) {
	g.Returns = val
}
