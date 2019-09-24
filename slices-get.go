package bean

type BoolSliceGetter struct {
	Calls   uint
	Returns []bool
}

type ByteSliceGetter struct {
	Calls   uint
	Returns []byte
}

type Complex64SliceGetter struct {
	Calls   uint
	Returns []complex64
}

type Complex128SliceGetter struct {
	Calls   uint
	Returns []complex128
}

type Float32SliceGetter struct {
	Calls   uint
	Returns []float32
}

type Float64SliceGetter struct {
	Calls   uint
	Returns []float64
}

type IntSliceGetter struct {
	Calls   uint
	Returns []int
}

type Int8SliceGetter struct {
	Calls   uint
	Returns []int8
}

type Int16SliceGetter struct {
	Calls   uint
	Returns []int16
}

type Int32SliceGetter struct {
	Calls   uint
	Returns []int32
}

type Int64SliceGetter struct {
	Calls   uint
	Returns []int64
}

type RuneSliceGetter struct {
	Calls   uint
	Returns []rune
}

type StringSliceGetter struct {
	Calls   uint
	Returns []string
}

type UintSliceGetter struct {
	Calls   uint
	Returns []uint
}

type Uint8SliceGetter struct {
	Calls   uint
	Returns []uint8
}

type Uint16SliceGetter struct {
	Calls   uint
	Returns []uint16
}

type Uint32SliceGetter struct {
	Calls   uint
	Returns []uint32
}

type Uint64SliceGetter struct {
	Calls   uint
	Returns []uint64
}

type UintptrSliceGetter struct {
	Calls   uint
	Returns []uintptr
}
