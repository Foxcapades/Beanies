// Generated at 2019-09-24T21:57:53-04:00
package bean

type StringSetter struct {
	Calls   uint
	Inputs  []string
}

func (g *StringSetter) Set(v string) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g StringSetter) CallCount() uint {
	return g.Calls
}

func (g *StringSetter) InputValues() []string {
	return g.Inputs
}
