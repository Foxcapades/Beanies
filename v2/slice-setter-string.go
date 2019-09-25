// Generated at 2019-09-24T21:57:54-04:00
package bean

type StringSliceSetter struct {
	Calls   uint
	Inputs  [][]string
}

func (g *StringSliceSetter) Set(v []string) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g StringSliceSetter) CallCount() uint {
	return g.Calls
}

func (g *StringSliceSetter) InputValues() [][]string {
	return g.Inputs
}
