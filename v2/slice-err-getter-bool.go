// Generated at 2019-09-24T21:57:34-04:00
package bean

type BoolSliceErrGetter struct {
	Calls   uint
	Returns []bool
	Error		error
}

func (g *BoolSliceErrGetter) Get() ([]bool, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g BoolSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *BoolSliceErrGetter) SetReturnValue(val []bool) {
	g.Returns = val
}

func (g *BoolSliceErrGetter) SetError(err error) {
	g.Error = err
}
