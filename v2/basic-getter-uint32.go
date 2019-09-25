// Generated at 2019-09-24T21:49:17-04:00
package bean

type Uint32Getter struct {
	Calls   uint
	Returns uint32
}

func (g *Uint32Getter) Get() uint32 {
	g.Calls++
	return g.Returns
}

func (g Uint32Getter) CallCount() uint {
	return g.Calls
}

func (g *Uint32Getter) SetReturnValue(val uint32) {
	g.Returns = val
}
