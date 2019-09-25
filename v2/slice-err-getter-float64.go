// Generated at 2019-09-24T21:57:42-04:00
package bean

type Float64SliceErrGetter struct {
	Calls   uint
	Returns []float64
	Error		error
}

func (g *Float64SliceErrGetter) Get() ([]float64, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Float64SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Float64SliceErrGetter) SetReturnValue(val []float64) {
	g.Returns = val
}

func (g *Float64SliceErrGetter) SetError(err error) {
	g.Error = err
}
