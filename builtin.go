package change

type Bool interface {
	Get() bool
	Set(v bool, chg Flags) Flags
}

type Uint8 interface {
	Get() uint8
	Set(v uint8, chg Flags) Flags
}

type Uint16 interface {
	Get() uint16
	Set(v uint16, chg Flags) Flags
}

type Uint32 interface {
	Get() uint32
	Set(v uint32, chg Flags) Flags
}

type Uint64 interface {
	Get() uint64
	Set(v uint64, chg Flags) Flags
}

type Int8 interface {
	Get() int8
	Set(v int8, chg Flags) Flags
}

type Int16 interface {
	Get() int16
	Set(v int16, chg Flags) Flags
}

type Int32 interface {
	Get() int32
	Set(v int32, chg Flags) Flags
}

type Int64 interface {
	Get() int64
	Set(v int64, chg Flags) Flags
}

type Float32 interface {
	Get() float32
	Set(v float32, chg Flags) Flags
}

type Float64 interface {
	Get() float64
	Set(v float64, chg Flags) Flags
}

type Complex64 interface {
	Get() complex64
	Set(v complex64, chg Flags) Flags
}

type Complex128 interface {
	Get() complex128
	Set(v complex128, chg Flags) Flags
}

type Byte interface {
	Get() byte
	Set(v byte, chg Flags) Flags
}

type Rune interface {
	Get() rune
	Set(v rune, chg Flags) Flags
}

type Uint interface {
	Get() uint
	Set(v uint, chg Flags) Flags
}

type Int interface {
	Get() int
	Set(v int, chg Flags) Flags
}

type UintPtr interface {
	Get() uintptr
	Set(v uintptr, chg Flags) Flags
}

type String interface {
	Get() string
	Set(v string, chg Flags) Flags
}
