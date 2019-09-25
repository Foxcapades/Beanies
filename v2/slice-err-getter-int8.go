// Generated at 2019-09-24T21:49:03-04:00
package bean

type Int8SliceErrGetter struct {
	Calls   uint
	Returns []int8
}

func (g *Int8SliceErrGetter) Get() []int8 {
	g.Calls++
	return g.Returns
}

func (g Int8SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int8SliceErrGetter) SetReturnValue(val []int8) {
	g.Returns = val
}
