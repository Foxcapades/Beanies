// Generated at 2019-09-24T21:57:47-04:00
package bean

type Int16SliceGetter struct {
	Calls   uint
	Returns []int16
}

func (g *Int16SliceGetter) Get() []int16 {
	g.Calls++
	return g.Returns
}

func (g Int16SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Int16SliceGetter) SetReturnValue(val []int16) {
	g.Returns = val
}
