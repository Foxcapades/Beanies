// Generated at 2019-09-24T21:57:41-04:00
package bean

type Float64Getter struct {
	Calls   uint
	Returns float64
}

func (g *Float64Getter) Get() float64 {
	g.Calls++
	return g.Returns
}

func (g Float64Getter) CallCount() uint {
	return g.Calls
}

func (g *Float64Getter) SetReturnValue(val float64) {
	g.Returns = val
}
