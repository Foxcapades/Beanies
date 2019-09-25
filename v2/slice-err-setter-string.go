// Generated at 2019-09-24T21:57:54-04:00
package bean

type StringSliceErrSetter struct {
	Calls   uint
	Inputs  [][]string
	Error   error
}

func (g *StringSliceErrSetter) Set(v []string) error {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
	return g.Error
}

func (g StringSliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *StringSliceErrSetter) InputValues() [][]string {
	return g.Inputs
}

func (g *StringSliceErrSetter) SetError(err error) {
	g.Error = err
}
