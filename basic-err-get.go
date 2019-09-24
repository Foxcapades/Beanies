package bean

type ByteErrGetter struct {
	Calls   uint
	Returns byte
	Error   error
}

type Complex64ErrGetter struct {
	Calls   uint
	Returns complex64
	Error   error
}

type Complex128ErrGetter struct {
	Calls   uint
	Returns complex128
	Error   error
}

type Float32ErrGetter struct {
	Calls   uint
	Returns float32
	Error   error
}

type Float64ErrGetter struct {
	Calls   uint
	Returns float64
	Error   error
}

type IntErrGetter struct {
	Calls   uint
	Returns int
	Error   error
}

type Int8ErrGetter struct {
	Calls   uint
	Returns int8
	Error   error
}

type Int16ErrGetter struct {
	Calls   uint
	Returns int16
	Error   error
}

type Int32ErrGetter struct {
	Calls   uint
	Returns int32
	Error   error
}

type Int64ErrGetter struct {
	Calls   uint
	Returns int64
	Error   error
}

type RuneErrGetter struct {
	Calls   uint
	Returns rune
	Error   error
}

type StringErrGetter struct {
	Calls   uint
	Returns string
	Error   error
}

type UintErrGetter struct {
	Calls   uint
	Returns uint
	Error   error
}

type Uint8ErrGetter struct {
	Calls   uint
	Returns uint8
	Error   error
}

type Uint16ErrGetter struct {
	Calls   uint
	Returns uint16
	Error   error
}

type Uint32ErrGetter struct {
	Calls   uint
	Returns uint32
	Error   error
}

type Uint64ErrGetter struct {
	Calls   uint
	Returns uint64
	Error   error
}

type UintptrErrGetter struct {
	Calls   uint
	Returns uintptr
	Error   error
}
