// Generated at 2019-09-24T21:58:02-04:00
package bean

type UintptrGetter struct {
	Calls   uint
	Returns uintptr
}

func (g *UintptrGetter) Get() uintptr {
	g.Calls++
	return g.Returns
}

func (g UintptrGetter) CallCount() uint {
	return g.Calls
}

func (g *UintptrGetter) SetReturnValue(val uintptr) {
	g.Returns = val
}
