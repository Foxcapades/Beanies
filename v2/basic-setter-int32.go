// Generated at 2019-09-24T21:49:05-04:00
package bean

type Int32Setter struct {
	Calls   uint
	Inputs  []int32
}

func (g *Int32Setter) Set(v int32) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Int32Setter) CallCount() uint {
	return g.Calls
}

func (g *Int32Setter) InputValues() []int32 {
	return g.Inputs
}
