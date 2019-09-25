// Generated at 2019-09-24T21:49:18-04:00
package bean

type Uint64Getter struct {
	Calls   uint
	Returns uint64
}

func (g *Uint64Getter) Get() uint64 {
	g.Calls++
	return g.Returns
}

func (g Uint64Getter) CallCount() uint {
	return g.Calls
}

func (g *Uint64Getter) SetReturnValue(val uint64) {
	g.Returns = val
}
