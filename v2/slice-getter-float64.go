// Generated at 2019-09-24T21:49:00-04:00
package bean

type Float64SliceGetter struct {
	Calls   uint
	Returns []float64
}

func (g *Float64SliceGetter) Get() []float64 {
	g.Calls++
	return g.Returns
}

func (g Float64SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Float64SliceGetter) SetReturnValue(val []float64) {
	g.Returns = val
}
