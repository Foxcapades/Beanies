// Generated at 2019-09-24T21:49:18-04:00
package bean

type Uint32SliceErrGetter struct {
	Calls   uint
	Returns []uint32
}

func (g *Uint32SliceErrGetter) Get() []uint32 {
	g.Calls++
	return g.Returns
}

func (g Uint32SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint32SliceErrGetter) SetReturnValue(val []uint32) {
	g.Returns = val
}
