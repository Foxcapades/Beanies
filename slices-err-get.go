package bean

type ByteSliceErrGetter struct {
	Calls   uint
	Returns []byte
	Error   error
}

type Complex64SliceErrGetter struct {
	Calls   uint
	Returns []complex64
	Error   error
}

type Complex128SliceErrGetter struct {
	Calls   uint
	Returns []complex128
	Error   error
}

type Float32SliceErrGetter struct {
	Calls   uint
	Returns []float32
	Error   error
}

type Float64SliceErrGetter struct {
	Calls   uint
	Returns []float64
	Error   error
}

type IntSliceErrGetter struct {
	Calls   uint
	Returns []int
	Error   error
}

type Int8SliceErrGetter struct {
	Calls   uint
	Returns []int8
	Error   error
}

type Int16SliceErrGetter struct {
	Calls   uint
	Returns []int16
	Error   error
}

type Int32SliceErrGetter struct {
	Calls   uint
	Returns []int32
	Error   error
}

type Int64SliceErrGetter struct {
	Calls   uint
	Returns []int64
	Error   error
}

type RuneSliceErrGetter struct {
	Calls   uint
	Returns []rune
	Error   error
}

type StringSliceErrGetter struct {
	Calls   uint
	Returns []string
	Error   error
}

type UintSliceErrGetter struct {
	Calls   uint
	Returns []uint
	Error   error
}

type Uint8SliceErrGetter struct {
	Calls   uint
	Returns []uint8
	Error   error
}

type Uint16SliceErrGetter struct {
	Calls   uint
	Returns []uint16
	Error   error
}

type Uint32SliceErrGetter struct {
	Calls   uint
	Returns []uint32
	Error   error
}

type Uint64SliceErrGetter struct {
	Calls   uint
	Returns []uint64
	Error   error
}

type UintptrSliceErrGetter struct {
	Calls   uint
	Returns []uintptr
	Error   error
}
