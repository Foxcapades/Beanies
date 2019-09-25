// Generated at 2019-09-24T21:57:46-04:00
package bean

type Int16Getter struct {
	Calls   uint
	Returns int16
}

func (g *Int16Getter) Get() int16 {
	g.Calls++
	return g.Returns
}

func (g Int16Getter) CallCount() uint {
	return g.Calls
}

func (g *Int16Getter) SetReturnValue(val int16) {
	g.Returns = val
}
