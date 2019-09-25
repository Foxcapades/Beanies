// Generated at 2019-09-24T21:57:36-04:00
package bean

type Complex64Getter struct {
	Calls   uint
	Returns complex64
}

func (g *Complex64Getter) Get() complex64 {
	g.Calls++
	return g.Returns
}

func (g Complex64Getter) CallCount() uint {
	return g.Calls
}

func (g *Complex64Getter) SetReturnValue(val complex64) {
	g.Returns = val
}
