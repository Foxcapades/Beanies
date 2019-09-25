// Generated at 2019-09-24T21:48:51-04:00
package bean

type BoolSetter struct {
	Calls   uint
	Inputs  []bool
}

func (g *BoolSetter) Set(v bool) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g BoolSetter) CallCount() uint {
	return g.Calls
}

func (g *BoolSetter) InputValues() []bool {
	return g.Inputs
}
