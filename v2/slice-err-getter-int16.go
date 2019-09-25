// Generated at 2019-09-24T21:57:47-04:00
package bean

type Int16SliceErrGetter struct {
	Calls   uint
	Returns []int16
	Error		error
}

func (g *Int16SliceErrGetter) Get() ([]int16, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Int16SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int16SliceErrGetter) SetReturnValue(val []int16) {
	g.Returns = val
}

func (g *Int16SliceErrGetter) SetError(err error) {
	g.Error = err
}
