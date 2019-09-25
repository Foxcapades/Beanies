// Generated at 2019-09-24T21:57:43-04:00
package bean

type IntSetter struct {
	Calls   uint
	Inputs  []int
}

func (g *IntSetter) Set(v int) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g IntSetter) CallCount() uint {
	return g.Calls
}

func (g *IntSetter) InputValues() []int {
	return g.Inputs
}
