// Generated at 2019-09-24T21:48:58-04:00
package bean

type Float32ErrGetter struct {
	Calls   uint
	Returns float32
	Error		error
}

func (g *Float32ErrGetter) Get() (float32, error) {
	g.Calls++
	return g.Returns, g.Error
}

func (g Float32ErrGetter) CallCount() uint {
	return g.Calls
}

func (g *Float32ErrGetter) SetReturnValue(val float32) {
	g.Returns = val
}

func (g *Float32ErrGetter) SetError(err error) {
	g.Error = err
}
