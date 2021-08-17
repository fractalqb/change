// Copyright 2021 Marcus Perlick
// This file is part of Go module github.com/fractalqb/change
//
// Foobar is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// Foobar is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.

package chgv

import "github.com/fractalqb/change"

// Bool tracks changes of bool values with caller-provided Flags.
// See also chgv package documentation.
type Bool bool

func (cv Bool) Get() bool { return bool(cv) }

func (cv *Bool) Set(v bool, chg change.Flags) change.Flags {
	if bool(*cv) == v {
		return 0
	}
	*cv = Bool(v)
	return chg
}

// Uint8 tracks changes of uint8 values with caller-provided Flags.
// See also chgv package documentation.
type Uint8 uint8

func (cv Uint8) Get() uint8 { return uint8(cv) }

func (cv *Uint8) Set(v uint8, chg change.Flags) change.Flags {
	if uint8(*cv) == v {
		return 0
	}
	*cv = Uint8(v)
	return chg
}

// Uint16 tracks changes of uint16 values with caller-provided Flags.
// See also chgv package documentation.
type Uint16 uint16

func (cv Uint16) Get() uint16 { return uint16(cv) }

func (cv *Uint16) Set(v uint16, chg change.Flags) change.Flags {
	if uint16(*cv) == v {
		return 0
	}
	*cv = Uint16(v)
	return chg
}

// Uint32 tracks changes of uint32 values with caller-provided Flags.
// See also chgv package documentation.
type Uint32 uint32

func (cv Uint32) Get() uint32 { return uint32(cv) }

func (cv *Uint32) Set(v uint32, chg change.Flags) change.Flags {
	if uint32(*cv) == v {
		return 0
	}
	*cv = Uint32(v)
	return chg
}

// Uint64 tracks changes of uint64 values with caller-provided Flags.
// See also chgv package documentation.
type Uint64 uint64

func (cv Uint64) Get() uint64 { return uint64(cv) }

func (cv *Uint64) Set(v uint64, chg change.Flags) change.Flags {
	if uint64(*cv) == v {
		return 0
	}
	*cv = Uint64(v)
	return chg
}

// Int8 tracks changes of int8 values with caller-provided Flags.
// See also chgv package documentation.
type Int8 int8

func (cv Int8) Get() int8 { return int8(cv) }

func (cv *Int8) Set(v int8, chg change.Flags) change.Flags {
	if int8(*cv) == v {
		return 0
	}
	*cv = Int8(v)
	return chg
}

// Int16 tracks changes of int16 values with caller-provided Flags.
// See also chgv package documentation.
type Int16 int16

func (cv Int16) Get() int16 { return int16(cv) }

func (cv *Int16) Set(v int16, chg change.Flags) change.Flags {
	if int16(*cv) == v {
		return 0
	}
	*cv = Int16(v)
	return chg
}

// Int32 tracks changes of int32 values with caller-provided Flags.
// See also chgv package documentation.
type Int32 int32

func (cv Int32) Get() int32 { return int32(cv) }

func (cv *Int32) Set(v int32, chg change.Flags) change.Flags {
	if int32(*cv) == v {
		return 0
	}
	*cv = Int32(v)
	return chg
}

// Int64 tracks changes of int64 values with caller-provided Flags.
// See also chgv package documentation.
type Int64 int64

func (cv Int64) Get() int64 { return int64(cv) }

func (cv *Int64) Set(v int64, chg change.Flags) change.Flags {
	if int64(*cv) == v {
		return 0
	}
	*cv = Int64(v)
	return chg
}

// Float32 tracks changes of float32 values with caller-provided Flags.
// See also chgv package documentation.
type Float32 float32

func (cv Float32) Get() float32 { return float32(cv) }

func (cv *Float32) Set(v float32, chg change.Flags) change.Flags {
	if float32(*cv) == v {
		return 0
	}
	*cv = Float32(v)
	return chg
}

// Float64 tracks changes of float64 values with caller-provided Flags.
// See also chgv package documentation.
type Float64 float64

func (cv Float64) Get() float64 { return float64(cv) }

func (cv *Float64) Set(v float64, chg change.Flags) change.Flags {
	if float64(*cv) == v {
		return 0
	}
	*cv = Float64(v)
	return chg
}

// Complex64 tracks changes of complex64 values with caller-provided Flags.
// See also chgv package documentation.
type Complex64 complex64

func (cv Complex64) Get() complex64 { return complex64(cv) }

func (cv *Complex64) Set(v complex64, chg change.Flags) change.Flags {
	if complex64(*cv) == v {
		return 0
	}
	*cv = Complex64(v)
	return chg
}

// Complex128 tracks changes of complex128 values with caller-provided Flags.
// See also chgv package documentation.
type Complex128 complex128

func (cv Complex128) Get() complex128 { return complex128(cv) }

func (cv *Complex128) Set(v complex128, chg change.Flags) change.Flags {
	if complex128(*cv) == v {
		return 0
	}
	*cv = Complex128(v)
	return chg
}

// Byte tracks changes of byte values with caller-provided Flags.
// See also chgv package documentation.
type Byte byte

func (cv Byte) Get() byte { return byte(cv) }

func (cv *Byte) Set(v byte, chg change.Flags) change.Flags {
	if byte(*cv) == v {
		return 0
	}
	*cv = Byte(v)
	return chg
}

// Rune tracks changes of rune values with caller-provided Flags.
// See also chgv package documentation.
type Rune rune

func (cv Rune) Get() rune { return rune(cv) }

func (cv *Rune) Set(v rune, chg change.Flags) change.Flags {
	if rune(*cv) == v {
		return 0
	}
	*cv = Rune(v)
	return chg
}

// Uint tracks changes of uint values with caller-provided Flags.
// See also chgv package documentation.
type Uint uint

func (cv Uint) Get() uint { return uint(cv) }

func (cv *Uint) Set(v uint, chg change.Flags) change.Flags {
	if uint(*cv) == v {
		return 0
	}
	*cv = Uint(v)
	return chg
}

// Int tracks changes of int values with caller-provided Flags.
// See also chgv package documentation.
type Int int

func (cv Int) Get() int { return int(cv) }

func (cv *Int) Set(v int, chg change.Flags) change.Flags {
	if int(*cv) == v {
		return 0
	}
	*cv = Int(v)
	return chg
}

// UintPtr tracks changes of uintptr values with caller-provided Flags.
// See also chgv package documentation.
type UintPtr uintptr

func (cv UintPtr) Get() uintptr { return uintptr(cv) }

func (cv *UintPtr) Set(v uintptr, chg change.Flags) change.Flags {
	if uintptr(*cv) == v {
		return 0
	}
	*cv = UintPtr(v)
	return chg
}

// String tracks changes of string values with caller-provided Flags.
// See also chgv package documentation.
type String string

func (cv String) Get() string { return string(cv) }

func (cv *String) Set(v string, chg change.Flags) change.Flags {
	if string(*cv) == v {
		return 0
	}
	*cv = String(v)
	return chg
}
