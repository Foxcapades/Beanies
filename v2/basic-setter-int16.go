// Generated at 2019-09-24T21:57:46-04:00
package bean

type Int16Setter struct {
	Calls   uint
	Inputs  []int16
}

func (g *Int16Setter) Set(v int16) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int16Setter) CallCount() uint {
	return g.Calls
}

func (g *Int16Setter) InputValues() []int16 {
	return g.Inputs
}
