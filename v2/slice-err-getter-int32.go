// Generated at 2019-09-24T21:57:49-04:00
package bean

type Int32SliceErrGetter struct {
	Calls   uint
	Returns []int32
	Error		error
}

func (g *Int32SliceErrGetter) Get() ([]int32, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Int32SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int32SliceErrGetter) SetReturnValue(val []int32) {
	g.Returns = val
}

func (g *Int32SliceErrGetter) SetError(err error) {
	g.Error = err
}
