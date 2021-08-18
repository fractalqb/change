define(`onchfalgfunc', `// $2 returns chg as change.Flags
func (chg ChgFlag) $2(src *$2, _, _ $1, _ bool) change.Flags {
	return change.Flags(chg)
}
')dnl
define(`onchgtype', `// $2 tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type $2 struct {
	v chgv.$2
	h $2Hook
}

// New$2 creates a change-detectable $2. See also section Hook
// Funtions.
func New$2(init $1, hook $2Hook) *$2 {
	return &$2{v: chgv.$2(init), h: hook}
}

// See section Hook Funtions.
type $2Hook func(src *$2, ov, nv $1, check bool) change.Flags

// Reset does not call a hook, even if the value changes.
// See also section Hook Funtions.
func (cv *$2) Reset(v $1, hook $2Hook) ($1, $2Hook) {
	ov, oh := cv.v, cv.h
	(*cv) = $2{v: chgv.$2(v), h: hook}
	return ov.Get(), oh
}

// Get returns the current $1 value.
func (cv $2) Get() $1 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
func (cv *$2) Set(v $1, chg change.Flags) change.Flags {
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
')dnl
dnl
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

// ChgFlag is used to provide hooks that always return a Flags
// value. It is used to provide default Flags when the caller of a Set
// operation passes 0 Flags. Each method of ChgFlag is a type specific
// hook for a change-detectable value in package onchg.
type ChgFlag change.Flags

onchfalgfunc(`bool', `Bool')
onchfalgfunc(`uint8', `Uint8')
onchfalgfunc(`uint16', `Uint16')
onchfalgfunc(`uint32', `Uint32')
onchfalgfunc(`uint64', `Uint64')
onchfalgfunc(`int8', `Int8')
onchfalgfunc(`int16', `Int16')
onchfalgfunc(`int32', `Int32')
onchfalgfunc(`int64', `Int64')
onchfalgfunc(`float32', `Float32')
onchfalgfunc(`float64', `Float64')
onchfalgfunc(`complex64', `Complex64')
onchfalgfunc(`complex128', `Complex128')
onchfalgfunc(`byte', `Byte')
onchfalgfunc(`rune', `Rune')
onchfalgfunc(`uint', `Uint')
onchfalgfunc(`int', `Int')
onchfalgfunc(`uintptr', `UintPtr')
onchfalgfunc(`string', `String')
dnl
onchgtype(`bool', `Bool')
onchgtype(`uint8', `Uint8')
onchgtype(`uint16', `Uint16')
onchgtype(`uint32', `Uint32')
onchgtype(`uint64', `Uint64')
onchgtype(`int8', `Int8')
onchgtype(`int16', `Int16')
onchgtype(`int32', `Int32')
onchgtype(`int64', `Int64')
onchgtype(`float32', `Float32')
onchgtype(`float64', `Float64')
onchgtype(`complex64', `Complex64')
onchgtype(`complex128', `Complex128')
onchgtype(`byte', `Byte')
onchgtype(`rune', `Rune')
onchgtype(`uint', `Uint')
onchgtype(`int', `Int')
onchgtype(`uintptr', `UintPtr')
onchgtype(`string', `String')
