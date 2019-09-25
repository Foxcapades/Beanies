// Generated at 2019-09-24T21:49:12-04:00
package bean

type UintSliceGetter struct {
	Calls   uint
	Returns []uint
}

func (g *UintSliceGetter) Get() []uint {
	g.Calls++
	return g.Returns
}

func (g UintSliceGetter) CallCount() uint {
	return g.Calls
}

func (g *UintSliceGetter) SetReturnValue(val []uint) {
	g.Returns = val
}
