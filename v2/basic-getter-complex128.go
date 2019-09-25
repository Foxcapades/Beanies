// Generated at 2019-09-24T21:48:56-04:00
package bean

type Complex128Getter struct {
	Calls   uint
	Returns complex128
}

func (g *Complex128Getter) Get() complex128 {
	g.Calls++
	return g.Returns
}

func (g Complex128Getter) CallCount() uint {
	return g.Calls
}

func (g *Complex128Getter) SetReturnValue(val complex128) {
	g.Returns = val
}
