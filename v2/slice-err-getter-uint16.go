// Generated at 2019-09-24T21:49:16-04:00
package bean

type Uint16SliceErrGetter struct {
	Calls   uint
	Returns []uint16
}

func (g *Uint16SliceErrGetter) Get() []uint16 {
	g.Calls++
	return g.Returns
}

func (g Uint16SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint16SliceErrGetter) SetReturnValue(val []uint16) {
	g.Returns = val
}