// Generated at 2019-09-24T21:58:00-04:00
package bean

type Uint32SliceErrGetter struct {
	Calls   uint
	Returns []uint32
	Error		error
}

func (g *Uint32SliceErrGetter) Get() ([]uint32, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Uint32SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint32SliceErrGetter) SetReturnValue(val []uint32) {
	g.Returns = val
}

func (g *Uint32SliceErrGetter) SetError(err error) {
	g.Error = err
}
