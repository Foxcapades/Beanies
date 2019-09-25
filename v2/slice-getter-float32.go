// Generated at 2019-09-24T21:48:58-04:00
package bean

type Float32SliceGetter struct {
	Calls   uint
	Returns []float32
}

func (g *Float32SliceGetter) Get() []float32 {
	g.Calls++
	return g.Returns
}

func (g Float32SliceGetter) CallCount() uint {
	return g.Calls
}

func (g *Float32SliceGetter) SetReturnValue(val []float32) {
	g.Returns = val
}
