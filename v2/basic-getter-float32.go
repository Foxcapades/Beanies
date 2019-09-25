// Generated at 2019-09-24T21:48:57-04:00
package bean

type Float32Getter struct {
	Calls   uint
	Returns float32
}

func (g *Float32Getter) Get() float32 {
	g.Calls++
	return g.Returns
}

func (g Float32Getter) CallCount() uint {
	return g.Calls
}

func (g *Float32Getter) SetReturnValue(val float32) {
	g.Returns = val
}
