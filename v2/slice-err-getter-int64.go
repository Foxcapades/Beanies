// Generated at 2019-09-24T21:49:08-04:00
package bean

type Int64SliceErrGetter struct {
	Calls   uint
	Returns []int64
}

func (g *Int64SliceErrGetter) Get() []int64 {
	g.Calls++
	return g.Returns
}

func (g Int64SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int64SliceErrGetter) SetReturnValue(val []int64) {
	g.Returns = val
}
