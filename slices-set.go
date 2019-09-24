package bean

type ByteSliceSetter struct {
	Calls uint
	Input [][]byte
}

type Complex64SliceSetter struct {
	Calls uint
	Input [][]complex64
}

type Complex128SliceSetter struct {
	Calls uint
	Input [][]complex128
}

type Float32SliceSetter struct {
	Calls uint
	Input [][]float32
}

type Float64SliceSetter struct {
	Calls uint
	Input [][]float64
}

type IntSliceSetter struct {
	Calls uint
	Input [][]int
}

type Int8SliceSetter struct {
	Calls uint
	Input [][]int8
}

type Int16SliceSetter struct {
	Calls uint
	Input [][]int16
}

type Int32SliceSetter struct {
	Calls uint
	Input [][]int32
}

type Int64SliceSetter struct {
	Calls uint
	Input [][]int64
}

type RuneSliceSetter struct {
	Calls uint
	Input [][]rune
}

type StringSliceSetter struct {
	Calls uint
	Input [][]string
}

type UintSliceSetter struct {
	Calls uint
	Input [][]uint
}

type Uint8SliceSetter struct {
	Calls uint
	Input [][]uint8
}

type Uint16SliceSetter struct {
	Calls uint
	Input [][]uint16
}

type Uint32SliceSetter struct {
	Calls uint
	Input [][]uint32
}

type Uint64SliceSetter struct {
	Calls uint
	Input [][]uint64
}

type UintptrSliceSetter struct {
	Calls uint
	Input [][]uintptr
}
