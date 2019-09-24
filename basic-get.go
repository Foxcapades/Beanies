package bean

type BoolGetter struct {
	Calls   uint
	Returns bool
}

type ByteGetter struct {
	Calls   uint
	Returns byte
}

type Complex64Getter struct {
	Calls   uint
	Returns complex64
}

type Complex128Getter struct {
	Calls   uint
	Returns complex128
}

type Float32Getter struct {
	Calls   uint
	Returns float32
}

type Float64Getter struct {
	Calls   uint
	Returns float64
}

type IntGetter struct {
	Calls   uint
	Returns int
}

type Int8Getter struct {
	Calls   uint
	Returns int8
}

type Int16Getter struct {
	Calls   uint
	Returns int16
}

type Int32Getter struct {
	Calls   uint
	Returns int32
}

type Int64Getter struct {
	Calls   uint
	Returns int64
}

type RuneGetter struct {
	Calls   uint
	Returns rune
}

type StringGetter struct {
	Calls   uint
	Returns string
}

type UintGetter struct {
	Calls   uint
	Returns uint
}

type Uint8Getter struct {
	Calls   uint
	Returns uint8
}

type Uint16Getter struct {
	Calls   uint
	Returns uint16
}

type Uint32Getter struct {
	Calls   uint
	Returns uint32
}

type Uint64Getter struct {
	Calls   uint
	Returns uint64
}

type UintptrGetter struct {
	Calls   uint
	Returns uintptr
}
