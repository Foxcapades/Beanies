// Generated at 2019-09-24T21:58:03-04:00
package bean

type UintptrSliceGetter struct {
	Calls   uint
	Returns []uintptr
}

func (g *UintptrSliceGetter) Get() []uintptr {
	g.Calls++
	return g.Returns
}

func (g UintptrSliceGetter) CallCount() uint {
	return g.Calls
}

func (g *UintptrSliceGetter) SetReturnValue(val []uintptr) {
	g.Returns = val
}
