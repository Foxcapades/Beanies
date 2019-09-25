// Generated at 2019-09-24T21:58:00-04:00
package bean

type Uint32SliceGetter struct {
	Calls   uint
	Returns []uint32
}

func (g *Uint32SliceGetter) Get() []uint32 {
	g.Calls++
	return g.Returns
}

func (g Uint32SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint32SliceGetter) SetReturnValue(val []uint32) {
	g.Returns = val
}
