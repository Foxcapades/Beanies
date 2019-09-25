// Generated at 2019-09-24T21:49:13-04:00
package bean

type UintSliceErrGetter struct {
	Calls   uint
	Returns []uint
}

func (g *UintSliceErrGetter) Get() []uint {
	g.Calls++
	return g.Returns
}

func (g UintSliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *UintSliceErrGetter) SetReturnValue(val []uint) {
	g.Returns = val
}
