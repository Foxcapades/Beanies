// Generated at 2019-09-24T21:57:54-04:00
package bean

type UintGetter struct {
	Calls   uint
	Returns uint
}

func (g *UintGetter) Get() uint {
	g.Calls++
	return g.Returns
}

func (g UintGetter) CallCount() uint {
	return g.Calls
}

func (g *UintGetter) SetReturnValue(val uint) {
	g.Returns = val
}
