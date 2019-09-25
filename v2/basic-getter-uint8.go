// Generated at 2019-09-24T21:49:13-04:00
package bean

type Uint8Getter struct {
	Calls   uint
	Returns uint8
}

func (g *Uint8Getter) Get() uint8 {
	g.Calls++
	return g.Returns
}

func (g Uint8Getter) CallCount() uint {
	return g.Calls
}

func (g *Uint8Getter) SetReturnValue(val uint8) {
	g.Returns = val
}
