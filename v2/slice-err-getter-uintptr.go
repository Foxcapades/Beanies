// Generated at 2019-09-24T21:49:21-04:00
package bean

type UintptrSliceErrGetter struct {
	Calls   uint
	Returns []uintptr
}

func (g *UintptrSliceErrGetter) Get() []uintptr {
	g.Calls++
	return g.Returns
}

func (g UintptrSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *UintptrSliceErrGetter) SetReturnValue(val []uintptr) {
	g.Returns = val
}
