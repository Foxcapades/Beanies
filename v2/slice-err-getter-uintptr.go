// Generated at 2019-09-24T21:58:04-04:00
package bean

type UintptrSliceErrGetter struct {
	Calls   uint
	Returns []uintptr
	Error		error
}

func (g *UintptrSliceErrGetter) Get() ([]uintptr, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g UintptrSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *UintptrSliceErrGetter) SetReturnValue(val []uintptr) {
	g.Returns = val
}

func (g *UintptrSliceErrGetter) SetError(err error) {
	g.Error = err
}
