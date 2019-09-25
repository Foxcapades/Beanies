// Generated at 2019-09-24T21:49:05-04:00
package bean

type Int32Getter struct {
	Calls   uint
	Returns int32
}

func (g *Int32Getter) Get() int32 {
	g.Calls++
	return g.Returns
}

func (g Int32Getter) CallCount() uint {
	return g.Calls
}

func (g *Int32Getter) SetReturnValue(val int32) {
	g.Returns = val
}
