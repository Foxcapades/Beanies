// Generated at 2019-09-24T21:49:14-04:00
package bean

type Uint8SliceErrGetter struct {
	Calls   uint
	Returns []uint8
}

func (g *Uint8SliceErrGetter) Get() []uint8 {
	g.Calls++
	return g.Returns
}

func (g Uint8SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint8SliceErrGetter) SetReturnValue(val []uint8) {
	g.Returns = val
}
