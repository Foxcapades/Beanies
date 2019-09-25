// Generated at 2019-09-24T21:57:44-04:00
package bean

type IntSliceSetter struct {
	Calls   uint
	Inputs  [][]int
}

func (g *IntSliceSetter) Set(v []int) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g IntSliceSetter) CallCount() uint {
	return g.Calls
}

func (g *IntSliceSetter) InputValues() [][]int {
	return g.Inputs
}
