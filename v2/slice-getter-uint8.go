// Generated at 2019-09-24T21:57:57-04:00
package bean

type Uint8SliceGetter struct {
	Calls   uint
	Returns []uint8
}

func (g *Uint8SliceGetter) Get() []uint8 {
	g.Calls++
	return g.Returns
}

func (g Uint8SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint8SliceGetter) SetReturnValue(val []uint8) {
	g.Returns = val
}
