package chgv

import "git.fractalqb.de/fractalqb/change"

type Bool bool

func (b Bool) Get() bool { return bool(b) }

func (b *Bool) Set(v bool, chg change.Flags) change.Flags {
	if bool(*b) == v {
		return 0
	}
	*b = Bool(v)
	return chg
}

type Uint8 uint8

func (i Uint8) Get() uint8 { return uint8(i) }

func (i *Uint8) Set(v uint8, chg change.Flags) change.Flags {
	if uint8(*i) == v {
		return 0
	}
	*i = Uint8(v)
	return chg
}

type Uint16 uint16

func (i Uint16) Get() uint16 { return uint16(i) }

func (i *Uint16) Set(v uint16, chg change.Flags) change.Flags {
	if uint16(*i) == v {
		return 0
	}
	*i = Uint16(v)
	return chg
}

type Uint32 uint32

func (i Uint32) Get() uint32 { return uint32(i) }

func (i *Uint32) Set(v uint32, chg change.Flags) change.Flags {
	if uint32(*i) == v {
		return 0
	}
	*i = Uint32(v)
	return chg
}

type Uint64 uint64

func (i Uint64) Get() uint64 { return uint64(i) }

func (i *Uint64) Set(v uint64, chg change.Flags) change.Flags {
	if uint64(*i) == v {
		return 0
	}
	*i = Uint64(v)
	return chg
}

type Int8 int8

func (i Int8) Get() int8 { return int8(i) }

func (i *Int8) Set(v int8, chg change.Flags) change.Flags {
	if int8(*i) == v {
		return 0
	}
	*i = Int8(v)
	return chg
}

type Int16 int16

func (i Int16) Get() int16 { return int16(i) }

func (i *Int16) Set(v int16, chg change.Flags) change.Flags {
	if int16(*i) == v {
		return 0
	}
	*i = Int16(v)
	return chg
}

type Int32 int32

func (i Int32) Get() int32 { return int32(i) }

func (i *Int32) Set(v int32, chg change.Flags) change.Flags {
	if int32(*i) == v {
		return 0
	}
	*i = Int32(v)
	return chg
}

type Int64 int64

func (i Int64) Get() int64 { return int64(i) }

func (i *Int64) Set(v int64, chg change.Flags) change.Flags {
	if int64(*i) == v {
		return 0
	}
	*i = Int64(v)
	return chg
}

type Float32 float32

func (f Float32) Get() float32 { return float32(f) }

func (f *Float32) Set(v float32, chg change.Flags) change.Flags {
	if float32(*f) == v {
		return 0
	}
	*f = Float32(v)
	return chg
}

type Float64 float64

func (f Float64) Get() float64 { return float64(f) }

func (f *Float64) Set(v float64, chg change.Flags) change.Flags {
	if float64(*f) == v {
		return 0
	}
	*f = Float64(v)
	return chg
}

type Complex64 complex64

func (f Complex64) Get() complex64 { return complex64(f) }

func (f *Complex64) Set(v complex64, chg change.Flags) change.Flags {
	if complex64(*f) == v {
		return 0
	}
	*f = Complex64(v)
	return chg
}

type Complex128 complex128

func (f Complex128) Get() complex128 { return complex128(f) }

func (f *Complex128) Set(v complex128, chg change.Flags) change.Flags {
	if complex128(*f) == v {
		return 0
	}
	*f = Complex128(v)
	return chg
}

type Byte byte

func (b Byte) Get() byte { return byte(b) }

func (b *Byte) Set(v byte, chg change.Flags) change.Flags {
	if byte(*b) == v {
		return 0
	}
	*b = Byte(v)
	return chg
}

type Rune rune

func (r Rune) Get() rune { return rune(r) }

func (r *Rune) Set(v rune, chg change.Flags) change.Flags {
	if rune(*r) == v {
		return 0
	}
	*r = Rune(v)
	return chg
}

type Uint uint

func (i Uint) Get() uint { return uint(i) }

func (i *Uint) Set(v uint, chg change.Flags) change.Flags {
	if uint(*i) == v {
		return 0
	}
	*i = Uint(v)
	return chg
}

type Int int

func (i Int) Get() int { return int(i) }

func (i *Int) Set(v int, chg change.Flags) change.Flags {
	if int(*i) == v {
		return 0
	}
	*i = Int(v)
	return chg
}

type UintPtr uintptr

func (i UintPtr) Get() uintptr { return uintptr(i) }

func (i *UintPtr) Set(v uintptr, chg change.Flags) change.Flags {
	if uintptr(*i) == v {
		return 0
	}
	*i = UintPtr(v)
	return chg
}

type String string

func (s String) Get() string { return string(s) }

func (s *String) Set(v string, chg change.Flags) change.Flags {
	if string(*s) == v {
		return 0
	}
	*s = String(v)
	return chg
}
