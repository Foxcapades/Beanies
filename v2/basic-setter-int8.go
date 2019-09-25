// Generated at 2019-09-24T21:49:02-04:00
package bean

type Int8Setter struct {
	Calls   uint
	Inputs  []int8
}

func (g *Int8Setter) Set(v int8) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int8Setter) CallCount() uint {
	return g.Calls
}

func (g *Int8Setter) InputValues() []int8 {
	return g.Inputs
}
