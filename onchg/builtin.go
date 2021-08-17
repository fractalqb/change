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

package onchg

import (
	"github.com/fractalqb/change"
	"github.com/fractalqb/change/chgv"
)

// ChgFlags is used to provide hooks that always return a Flags value. It is
// used to provide default Flags when the caller of a Set operation passes
// 0 Flags. Each method of ChgFlags is a type specific hook for a
// change-detectable value in package onchg.
type ChgFlag change.Flags

// Bool returns chg as change.Flags
func (chg ChgFlag) Bool(src *Bool, _, _ bool, _ bool) change.Flags {
	return change.Flags(chg)
}

// Uint8 returns chg as change.Flags
func (chg ChgFlag) Uint8(src *Uint8, _, _ uint8, _ bool) change.Flags {
	return change.Flags(chg)
}

// Uint16 returns chg as change.Flags
func (chg ChgFlag) Uint16(src *Uint16, _, _ uint16, _ bool) change.Flags {
	return change.Flags(chg)
}

// Uint32 returns chg as change.Flags
func (chg ChgFlag) Uint32(src *Uint32, _, _ uint32, _ bool) change.Flags {
	return change.Flags(chg)
}

// Uint64 returns chg as change.Flags
func (chg ChgFlag) Uint64(src *Uint64, _, _ uint64, _ bool) change.Flags {
	return change.Flags(chg)
}

// Int8 returns chg as change.Flags
func (chg ChgFlag) Int8(src *Int8, _, _ int8, _ bool) change.Flags {
	return change.Flags(chg)
}

// Int16 returns chg as change.Flags
func (chg ChgFlag) Int16(src *Int16, _, _ int16, _ bool) change.Flags {
	return change.Flags(chg)
}

// Int32 returns chg as change.Flags
func (chg ChgFlag) Int32(src *Int32, _, _ int32, _ bool) change.Flags {
	return change.Flags(chg)
}

// Int64 returns chg as change.Flags
func (chg ChgFlag) Int64(src *Int64, _, _ int64, _ bool) change.Flags {
	return change.Flags(chg)
}

// Float32 returns chg as change.Flags
func (chg ChgFlag) Float32(src *Float32, _, _ float32, _ bool) change.Flags {
	return change.Flags(chg)
}

// Float64 returns chg as change.Flags
func (chg ChgFlag) Float64(src *Float64, _, _ float64, _ bool) change.Flags {
	return change.Flags(chg)
}

// Complex64 returns chg as change.Flags
func (chg ChgFlag) Complex64(src *Complex64, _, _ complex64, _ bool) change.Flags {
	return change.Flags(chg)
}

// Complex128 returns chg as change.Flags
func (chg ChgFlag) Complex128(src *Complex128, _, _ complex128, _ bool) change.Flags {
	return change.Flags(chg)
}

// Byte returns chg as change.Flags
func (chg ChgFlag) Byte(src *Byte, _, _ byte, _ bool) change.Flags {
	return change.Flags(chg)
}

// Rune returns chg as change.Flags
func (chg ChgFlag) Rune(src *Rune, _, _ rune, _ bool) change.Flags {
	return change.Flags(chg)
}

// Uint returns chg as change.Flags
func (chg ChgFlag) Uint(src *Uint, _, _ uint, _ bool) change.Flags {
	return change.Flags(chg)
}

// Int returns chg as change.Flags
func (chg ChgFlag) Int(src *Int, _, _ int, _ bool) change.Flags {
	return change.Flags(chg)
}

// UintPtr returns chg as change.Flags
func (chg ChgFlag) UintPtr(src *UintPtr, _, _ uintptr, _ bool) change.Flags {
	return change.Flags(chg)
}

// String returns chg as change.Flags
func (chg ChgFlag) String(src *String, _, _ string, _ bool) change.Flags {
	return change.Flags(chg)
}

// Bool tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Bool struct {
	v chgv.Bool
	h BoolHook
}

func NewBool(init bool, hook BoolHook) *Bool {
	return &Bool{v: chgv.Bool(init), h: hook}
}

type BoolHook func(src *Bool, ov, nv bool, check bool) change.Flags

func (cv *Bool) Reset(v bool, hook BoolHook) (bool, BoolHook) {
	ov, oh := cv.v, cv.h
	(*cv) = Bool{v: chgv.Bool(v), h: hook}
	return ov.Get(), oh
}

func (cv Bool) Get() bool { return cv.v.Get() }

func (cv *Bool) Set(v bool, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Uint8 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Uint8 struct {
	v chgv.Uint8
	h Uint8Hook
}

func NewUint8(init uint8, hook Uint8Hook) *Uint8 {
	return &Uint8{v: chgv.Uint8(init), h: hook}
}

type Uint8Hook func(src *Uint8, ov, nv uint8, check bool) change.Flags

func (cv *Uint8) Reset(v uint8, hook Uint8Hook) (uint8, Uint8Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Uint8{v: chgv.Uint8(v), h: hook}
	return ov.Get(), oh
}

func (cv Uint8) Get() uint8 { return cv.v.Get() }

func (cv *Uint8) Set(v uint8, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Uint16 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Uint16 struct {
	v chgv.Uint16
	h Uint16Hook
}

func NewUint16(init uint16, hook Uint16Hook) *Uint16 {
	return &Uint16{v: chgv.Uint16(init), h: hook}
}

type Uint16Hook func(src *Uint16, ov, nv uint16, check bool) change.Flags

func (cv *Uint16) Reset(v uint16, hook Uint16Hook) (uint16, Uint16Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Uint16{v: chgv.Uint16(v), h: hook}
	return ov.Get(), oh
}

func (cv Uint16) Get() uint16 { return cv.v.Get() }

func (cv *Uint16) Set(v uint16, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Uint32 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Uint32 struct {
	v chgv.Uint32
	h Uint32Hook
}

func NewUint32(init uint32, hook Uint32Hook) *Uint32 {
	return &Uint32{v: chgv.Uint32(init), h: hook}
}

type Uint32Hook func(src *Uint32, ov, nv uint32, check bool) change.Flags

func (cv *Uint32) Reset(v uint32, hook Uint32Hook) (uint32, Uint32Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Uint32{v: chgv.Uint32(v), h: hook}
	return ov.Get(), oh
}

func (cv Uint32) Get() uint32 { return cv.v.Get() }

func (cv *Uint32) Set(v uint32, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Uint64 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Uint64 struct {
	v chgv.Uint64
	h Uint64Hook
}

func NewUint64(init uint64, hook Uint64Hook) *Uint64 {
	return &Uint64{v: chgv.Uint64(init), h: hook}
}

type Uint64Hook func(src *Uint64, ov, nv uint64, check bool) change.Flags

func (cv *Uint64) Reset(v uint64, hook Uint64Hook) (uint64, Uint64Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Uint64{v: chgv.Uint64(v), h: hook}
	return ov.Get(), oh
}

func (cv Uint64) Get() uint64 { return cv.v.Get() }

func (cv *Uint64) Set(v uint64, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Int8 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Int8 struct {
	v chgv.Int8
	h Int8Hook
}

func NewInt8(init int8, hook Int8Hook) *Int8 {
	return &Int8{v: chgv.Int8(init), h: hook}
}

type Int8Hook func(src *Int8, ov, nv int8, check bool) change.Flags

func (cv *Int8) Reset(v int8, hook Int8Hook) (int8, Int8Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Int8{v: chgv.Int8(v), h: hook}
	return ov.Get(), oh
}

func (cv Int8) Get() int8 { return cv.v.Get() }

func (cv *Int8) Set(v int8, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Int16 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Int16 struct {
	v chgv.Int16
	h Int16Hook
}

func NewInt16(init int16, hook Int16Hook) *Int16 {
	return &Int16{v: chgv.Int16(init), h: hook}
}

type Int16Hook func(src *Int16, ov, nv int16, check bool) change.Flags

func (cv *Int16) Reset(v int16, hook Int16Hook) (int16, Int16Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Int16{v: chgv.Int16(v), h: hook}
	return ov.Get(), oh
}

func (cv Int16) Get() int16 { return cv.v.Get() }

func (cv *Int16) Set(v int16, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Int32 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Int32 struct {
	v chgv.Int32
	h Int32Hook
}

func NewInt32(init int32, hook Int32Hook) *Int32 {
	return &Int32{v: chgv.Int32(init), h: hook}
}

type Int32Hook func(src *Int32, ov, nv int32, check bool) change.Flags

func (cv *Int32) Reset(v int32, hook Int32Hook) (int32, Int32Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Int32{v: chgv.Int32(v), h: hook}
	return ov.Get(), oh
}

func (cv Int32) Get() int32 { return cv.v.Get() }

func (cv *Int32) Set(v int32, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Int64 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Int64 struct {
	v chgv.Int64
	h Int64Hook
}

func NewInt64(init int64, hook Int64Hook) *Int64 {
	return &Int64{v: chgv.Int64(init), h: hook}
}

type Int64Hook func(src *Int64, ov, nv int64, check bool) change.Flags

func (cv *Int64) Reset(v int64, hook Int64Hook) (int64, Int64Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Int64{v: chgv.Int64(v), h: hook}
	return ov.Get(), oh
}

func (cv Int64) Get() int64 { return cv.v.Get() }

func (cv *Int64) Set(v int64, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Float32 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Float32 struct {
	v chgv.Float32
	h Float32Hook
}

func NewFloat32(init float32, hook Float32Hook) *Float32 {
	return &Float32{v: chgv.Float32(init), h: hook}
}

type Float32Hook func(src *Float32, ov, nv float32, check bool) change.Flags

func (cv *Float32) Reset(v float32, hook Float32Hook) (float32, Float32Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Float32{v: chgv.Float32(v), h: hook}
	return ov.Get(), oh
}

func (cv Float32) Get() float32 { return cv.v.Get() }

func (cv *Float32) Set(v float32, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Float64 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Float64 struct {
	v chgv.Float64
	h Float64Hook
}

func NewFloat64(init float64, hook Float64Hook) *Float64 {
	return &Float64{v: chgv.Float64(init), h: hook}
}

type Float64Hook func(src *Float64, ov, nv float64, check bool) change.Flags

func (cv *Float64) Reset(v float64, hook Float64Hook) (float64, Float64Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Float64{v: chgv.Float64(v), h: hook}
	return ov.Get(), oh
}

func (cv Float64) Get() float64 { return cv.v.Get() }

func (cv *Float64) Set(v float64, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Complex64 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Complex64 struct {
	v chgv.Complex64
	h Complex64Hook
}

func NewComplex64(init complex64, hook Complex64Hook) *Complex64 {
	return &Complex64{v: chgv.Complex64(init), h: hook}
}

type Complex64Hook func(src *Complex64, ov, nv complex64, check bool) change.Flags

func (cv *Complex64) Reset(v complex64, hook Complex64Hook) (complex64, Complex64Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Complex64{v: chgv.Complex64(v), h: hook}
	return ov.Get(), oh
}

func (cv Complex64) Get() complex64 { return cv.v.Get() }

func (cv *Complex64) Set(v complex64, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Complex128 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Complex128 struct {
	v chgv.Complex128
	h Complex128Hook
}

func NewComplex128(init complex128, hook Complex128Hook) *Complex128 {
	return &Complex128{v: chgv.Complex128(init), h: hook}
}

type Complex128Hook func(src *Complex128, ov, nv complex128, check bool) change.Flags

func (cv *Complex128) Reset(v complex128, hook Complex128Hook) (complex128, Complex128Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = Complex128{v: chgv.Complex128(v), h: hook}
	return ov.Get(), oh
}

func (cv Complex128) Get() complex128 { return cv.v.Get() }

func (cv *Complex128) Set(v complex128, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Byte tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Byte struct {
	v chgv.Byte
	h ByteHook
}

func NewByte(init byte, hook ByteHook) *Byte {
	return &Byte{v: chgv.Byte(init), h: hook}
}

type ByteHook func(src *Byte, ov, nv byte, check bool) change.Flags

func (cv *Byte) Reset(v byte, hook ByteHook) (byte, ByteHook) {
	ov, oh := cv.v, cv.h
	(*cv) = Byte{v: chgv.Byte(v), h: hook}
	return ov.Get(), oh
}

func (cv Byte) Get() byte { return cv.v.Get() }

func (cv *Byte) Set(v byte, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Rune tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Rune struct {
	v chgv.Rune
	h RuneHook
}

func NewRune(init rune, hook RuneHook) *Rune {
	return &Rune{v: chgv.Rune(init), h: hook}
}

type RuneHook func(src *Rune, ov, nv rune, check bool) change.Flags

func (cv *Rune) Reset(v rune, hook RuneHook) (rune, RuneHook) {
	ov, oh := cv.v, cv.h
	(*cv) = Rune{v: chgv.Rune(v), h: hook}
	return ov.Get(), oh
}

func (cv Rune) Get() rune { return cv.v.Get() }

func (cv *Rune) Set(v rune, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Uint tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Uint struct {
	v chgv.Uint
	h UintHook
}

func NewUint(init uint, hook UintHook) *Uint {
	return &Uint{v: chgv.Uint(init), h: hook}
}

type UintHook func(src *Uint, ov, nv uint, check bool) change.Flags

func (cv *Uint) Reset(v uint, hook UintHook) (uint, UintHook) {
	ov, oh := cv.v, cv.h
	(*cv) = Uint{v: chgv.Uint(v), h: hook}
	return ov.Get(), oh
}

func (cv Uint) Get() uint { return cv.v.Get() }

func (cv *Uint) Set(v uint, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// Int tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Int struct {
	v chgv.Int
	h IntHook
}

func NewInt(init int, hook IntHook) *Int {
	return &Int{v: chgv.Int(init), h: hook}
}

type IntHook func(src *Int, ov, nv int, check bool) change.Flags

func (cv *Int) Reset(v int, hook IntHook) (int, IntHook) {
	ov, oh := cv.v, cv.h
	(*cv) = Int{v: chgv.Int(v), h: hook}
	return ov.Get(), oh
}

func (cv Int) Get() int { return cv.v.Get() }

func (cv *Int) Set(v int, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// UintPtr tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type UintPtr struct {
	v chgv.UintPtr
	h UintPtrHook
}

func NewUintPtr(init uintptr, hook UintPtrHook) *UintPtr {
	return &UintPtr{v: chgv.UintPtr(init), h: hook}
}

type UintPtrHook func(src *UintPtr, ov, nv uintptr, check bool) change.Flags

func (cv *UintPtr) Reset(v uintptr, hook UintPtrHook) (uintptr, UintPtrHook) {
	ov, oh := cv.v, cv.h
	(*cv) = UintPtr{v: chgv.UintPtr(v), h: hook}
	return ov.Get(), oh
}

func (cv UintPtr) Get() uintptr { return cv.v.Get() }

func (cv *UintPtr) Set(v uintptr, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}

// String tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type String struct {
	v chgv.String
	h StringHook
}

func NewString(init string, hook StringHook) *String {
	return &String{v: chgv.String(init), h: hook}
}

type StringHook func(src *String, ov, nv string, check bool) change.Flags

func (cv *String) Reset(v string, hook StringHook) (string, StringHook) {
	ov, oh := cv.v, cv.h
	(*cv) = String{v: chgv.String(v), h: hook}
	return ov.Get(), oh
}

func (cv String) Get() string { return cv.v.Get() }

func (cv *String) Set(v string, chg change.Flags) change.Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	if c := cv.h(cv, cv.v.Get(), v, true); c == 0 {
		return 0
	} else if chg == 0 {
		chg = c
	}
	old := cv.v.Get()
	if chg = cv.v.Set(v, chg); chg != 0 {
		cv.h(cv, old, v, false)
	}
	return chg
}