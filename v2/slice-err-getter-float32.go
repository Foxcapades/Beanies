// Generated at 2019-09-24T21:57:41-04:00
package bean

type Float32SliceErrGetter struct {
	Calls   uint
	Returns []float32
	Error		error
}

func (g *Float32SliceErrGetter) Get() ([]float32, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Float32SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Float32SliceErrGetter) SetReturnValue(val []float32) {
	g.Returns = val
}

func (g *Float32SliceErrGetter) SetError(err error) {
	g.Error = err
}
