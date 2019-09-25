// Generated at 2019-09-24T21:57:58-04:00
package bean

type Uint16SliceGetter struct {
	Calls   uint
	Returns []uint16
}

func (g *Uint16SliceGetter) Get() []uint16 {
	g.Calls++
	return g.Returns
}

func (g Uint16SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint16SliceGetter) SetReturnValue(val []uint16) {
	g.Returns = val
}
