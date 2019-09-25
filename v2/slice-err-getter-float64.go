// Generated at 2019-09-24T21:49:00-04:00
package bean

type Float64SliceErrGetter struct {
	Calls   uint
	Returns []float64
}

func (g *Float64SliceErrGetter) Get() []float64 {
	g.Calls++
	return g.Returns
}

func (g Float64SliceErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Float64SliceErrGetter) SetReturnValue(val []float64) {
	g.Returns = val
}
