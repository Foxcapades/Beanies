// Generated at 2019-09-24T21:49:06-04:00
package bean

type Int32SliceErrGetter struct {
	Calls   uint
	Returns []int32
}

func (g *Int32SliceErrGetter) Get() []int32 {
	g.Calls++
	return g.Returns
}

func (g Int32SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int32SliceErrGetter) SetReturnValue(val []int32) {
	g.Returns = val
}
