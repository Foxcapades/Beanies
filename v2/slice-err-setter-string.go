// Generated at 2019-09-24T21:49:11-04:00
package bean

type StringSliceErrSetter struct {
	Calls   uint
	Inputs  [][]string
}

func (g *StringSliceErrSetter) Set(v []string) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g StringSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *StringSliceErrSetter) InputValues() [][]string {
	return g.Inputs
}
