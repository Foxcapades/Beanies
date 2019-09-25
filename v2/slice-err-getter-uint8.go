// Generated at 2019-09-24T21:57:57-04:00
package bean

type Uint8SliceErrGetter struct {
	Calls   uint
	Returns []uint8
	Error		error
}

func (g *Uint8SliceErrGetter) Get() ([]uint8, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Uint8SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Uint8SliceErrGetter) SetReturnValue(val []uint8) {
	g.Returns = val
}

func (g *Uint8SliceErrGetter) SetError(err error) {
	g.Error = err
}
