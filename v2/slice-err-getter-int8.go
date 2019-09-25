// Generated at 2019-09-24T21:57:46-04:00
package bean

type Int8SliceErrGetter struct {
	Calls   uint
	Returns []int8
	Error		error
}

func (g *Int8SliceErrGetter) Get() ([]int8, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Int8SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Int8SliceErrGetter) SetReturnValue(val []int8) {
	g.Returns = val
}

func (g *Int8SliceErrGetter) SetError(err error) {
	g.Error = err
}
