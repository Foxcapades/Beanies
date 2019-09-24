package bean

type BoolSetter struct {
	Calls uint
	Input []bool
}

type ByteSetter struct {
	Calls uint
	Input []byte
}

type Complex64Setter struct {
	Calls uint
	Input []complex64
}

type Complex128Setter struct {
	Calls uint
	Input []complex128
}

type Float32Setter struct {
	Calls uint
	Input []float32
}

type Float64Setter struct {
	Calls uint
	Input []float64
}

type IntSetter struct {
	Calls uint
	Input []int
}

type Int8Setter struct {
	Calls uint
	Input []int8
}

type Int16Setter struct {
	Calls uint
	Input []int16
}

type Int32Setter struct {
	Calls uint
	Input []int32
}

type Int64Setter struct {
	Calls uint
	Input []int64
}

type RuneSetter struct {
	Calls uint
	Input []rune
}

type StringSetter struct {
	Calls uint
	Input []string
}

type UintSetter struct {
	Calls uint
	Input []uint
}

type Uint8Setter struct {
	Calls uint
	Input []uint8
}

type Uint16Setter struct {
	Calls uint
	Input []uint16
}

type Uint32Setter struct {
	Calls uint
	Input []uint32
}

type Uint64Setter struct {
	Calls uint
	Input []uint64
}

type UintptrSetter struct {
	Calls uint
	Input []uintptr
}
