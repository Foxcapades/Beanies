// Generated at 2019-09-24T21:57:51-04:00
package bean

type Int64SliceErrGetter struct {
	Calls   uint
	Returns []int64
	Error		error
}

func (g *Int64SliceErrGetter) Get() ([]int64, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Int64SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int64SliceErrGetter) SetReturnValue(val []int64) {
	g.Returns = val
}

func (g *Int64SliceErrGetter) SetError(err error) {
	g.Error = err
}
