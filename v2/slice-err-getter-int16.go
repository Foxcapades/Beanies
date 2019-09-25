// Generated at 2019-09-24T21:49:05-04:00
package bean

type Int16SliceErrGetter struct {
	Calls   uint
	Returns []int16
}

func (g *Int16SliceErrGetter) Get() []int16 {
	g.Calls++
	return g.Returns
}

func (g Int16SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int16SliceErrGetter) SetReturnValue(val []int16) {
	g.Returns = val
}
