// Generated at 2019-09-24T21:57:59-04:00
package bean

type Uint16SliceErrGetter struct {
	Calls   uint
	Returns []uint16
	Error		error
}

func (g *Uint16SliceErrGetter) Get() ([]uint16, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Uint16SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint16SliceErrGetter) SetReturnValue(val []uint16) {
	g.Returns = val
}

func (g *Uint16SliceErrGetter) SetError(err error) {
	g.Error = err
}
