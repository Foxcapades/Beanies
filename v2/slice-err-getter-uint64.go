// Generated at 2019-09-24T21:58:02-04:00
package bean

type Uint64SliceErrGetter struct {
	Calls   uint
	Returns []uint64
	Error		error
}

func (g *Uint64SliceErrGetter) Get() ([]uint64, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Uint64SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint64SliceErrGetter) SetReturnValue(val []uint64) {
	g.Returns = val
}

func (g *Uint64SliceErrGetter) SetError(err error) {
	g.Error = err
}
