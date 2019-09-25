// Generated at 2019-09-24T21:49:16-04:00
package bean

type Uint16SliceErrSetter struct {
	Calls   uint
	Inputs  [][]uint16
}

func (g *Uint16SliceErrSetter) Set(v []uint16) {
	g.Calls++
	g.Inputs = append(g.Inputs, v)
}

func (g Uint16SliceErrSetter) CallCount() uint {
	return g.Calls
}

func (g *Uint16SliceErrSetter) InputValues() [][]uint16 {
	return g.Inputs
}
