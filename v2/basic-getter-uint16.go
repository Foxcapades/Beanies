// Generated at 2019-09-24T21:49:15-04:00
package bean

type Uint16Getter struct {
	Calls   uint
	Returns uint16
}

func (g *Uint16Getter) Get() uint16 {
	g.Calls++
	return g.Returns
}

func (g Uint16Getter) CallCount() uint {
	return g.Calls
}

func (g *Uint16Getter) SetReturnValue(val uint16) {
	g.Returns = val
}
