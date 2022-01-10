// Copyright 2021 Marcus Perlick
// This file is part of Go module github.com/fractalqb/change
//
// change is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// change is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with change.  If not, see <http://www.gnu.org/licenses/>.

package change

// Bool is the common interface of change-detecable bool values.
type Bool interface {
	Get() bool
	Set(v bool, chg Flags) Flags
}

// Uint8 is the common interface of change-detecable uint8 values.
type Uint8 interface {
	Get() uint8
	Set(v uint8, chg Flags) Flags
}

// Uint16 is the common interface of change-detecable uint16 values.
type Uint16 interface {
	Get() uint16
	Set(v uint16, chg Flags) Flags
}

// Uint32 is the common interface of change-detecable uint32 values.
type Uint32 interface {
	Get() uint32
	Set(v uint32, chg Flags) Flags
}

// Uint64 is the common interface of change-detecable uint64 values.
type Uint64 interface {
	Get() uint64
	Set(v uint64, chg Flags) Flags
}

// Int8 is the common interface of change-detecable int8 values.
type Int8 interface {
	Get() int8
	Set(v int8, chg Flags) Flags
}

// Int16 is the common interface of change-detecable int16 values.
type Int16 interface {
	Get() int16
	Set(v int16, chg Flags) Flags
}

// Int32 is the common interface of change-detecable int32 values.
type Int32 interface {
	Get() int32
	Set(v int32, chg Flags) Flags
}

// Int64 is the common interface of change-detecable int64 values.
type Int64 interface {
	Get() int64
	Set(v int64, chg Flags) Flags
}

// Float32 is the common interface of change-detecable float32 values.
type Float32 interface {
	Get() float32
	Set(v float32, chg Flags) Flags
}

// Float64 is the common interface of change-detecable float64 values.
type Float64 interface {
	Get() float64
	Set(v float64, chg Flags) Flags
}

// Complex64 is the common interface of change-detecable complex64 values.
type Complex64 interface {
	Get() complex64
	Set(v complex64, chg Flags) Flags
}

// Complex128 is the common interface of change-detecable complex128 values.
type Complex128 interface {
	Get() complex128
	Set(v complex128, chg Flags) Flags
}

// Byte is the common interface of change-detecable byte values.
type Byte interface {
	Get() byte
	Set(v byte, chg Flags) Flags
}

// Rune is the common interface of change-detecable rune values.
type Rune interface {
	Get() rune
	Set(v rune, chg Flags) Flags
}

// Uint is the common interface of change-detecable uint values.
type Uint interface {
	Get() uint
	Set(v uint, chg Flags) Flags
}

// Int is the common interface of change-detecable int values.
type Int interface {
	Get() int
	Set(v int, chg Flags) Flags
}

// UintPtr is the common interface of change-detecable uintptr values.
type UintPtr interface {
	Get() uintptr
	Set(v uintptr, chg Flags) Flags
}

// String is the common interface of change-detecable string values.
type String interface {
	Get() string
	Set(v string, chg Flags) Flags
}
