// Generated at 2019-09-24T21:49:02-04:00
package bean

type Int8Getter struct {
	Calls   uint
	Returns int8
}

func (g *Int8Getter) Get() int8 {
	g.Calls++
	return g.Returns
}

func (g Int8Getter) CallCount() uint {
	return g.Calls
}

func (g *Int8Getter) SetReturnValue(val int8) {
	g.Returns = val
}
