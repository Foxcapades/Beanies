// Generated at 2019-09-24T21:49:07-04:00
package bean

type Int64Getter struct {
	Calls   uint
	Returns int64
}

func (g *Int64Getter) Get() int64 {
	g.Calls++
	return g.Returns
}

func (g Int64Getter) CallCount() uint {
	return g.Calls
}

func (g *Int64Getter) SetReturnValue(val int64) {
	g.Returns = val
}
