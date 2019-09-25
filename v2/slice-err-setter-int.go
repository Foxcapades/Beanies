// Generated at 2019-09-24T21:49:02-04:00
package bean

type IntSliceErrSetter struct {
	Calls   uint
	Inputs  [][]int
}

func (g *IntSliceErrSetter) Set(v []int) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g IntSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *IntSliceErrSetter) InputValues() [][]int {
	return g.Inputs
}
