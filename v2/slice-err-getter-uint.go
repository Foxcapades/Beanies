// Generated at 2019-09-24T21:57:55-04:00
package bean

type UintSliceErrGetter struct {
	Calls   uint
	Returns []uint
	Error		error
}

func (g *UintSliceErrGetter) Get() ([]uint, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g UintSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *UintSliceErrGetter) SetReturnValue(val []uint) {
	g.Returns = val
}

func (g *UintSliceErrGetter) SetError(err error) {
	g.Error = err
}
