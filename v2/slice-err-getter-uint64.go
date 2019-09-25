// Generated at 2019-09-24T21:49:19-04:00
package bean

type Uint64SliceErrGetter struct {
	Calls   uint
	Returns []uint64
}

func (g *Uint64SliceErrGetter) Get() []uint64 {
	g.Calls++
	return g.Returns
}

func (g Uint64SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint64SliceErrGetter) SetReturnValue(val []uint64) {
	g.Returns = val
}
