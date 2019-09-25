// Generated at 2019-09-24T21:58:03-04:00
package bean

type UintptrErrGetter struct {
	Calls   uint
	Returns uintptr
	Error		error
}

func (g *UintptrErrGetter) Get() (uintptr, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g UintptrErrGetter) CallCount() uint {
	return g.Calls
}

func (g *UintptrErrGetter) SetReturnValue(val uintptr) {
	g.Returns = val
}

func (g *UintptrErrGetter) SetError(err error) {
	g.Error = err
}
