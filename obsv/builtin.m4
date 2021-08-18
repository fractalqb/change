changequote(`<[', `]>')dnl
define(<[onchgtype]>, <[// $2 is an Observable $1 value. $2 implements change.$2.
type $2 struct {
	v chgv.$2
	ObservableBase
}

// New$2 creates a change-detectable $2 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func New$2(init $1, defaultTag interface{}, defaultChg change.Flags) *$2 {
	res := &$2{v: chgv.$2(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current $1 value.
func (cv $2) Get() $1 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
func (cv *$2) Set(v $1, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := $2Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.update(e)
	return chg
}

// $2Chg is the value change Event for $2.
type $2Chg struct {
	eventBase
	ov, nv $1
}

func (ce $2Chg) String() string {
	return fmt.Sprintf("Chg $1 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce $2Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce $2Chg) NewValue() interface{} { return ce.nv }
]>)dnl
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

package obsv

import (
	"fmt"

	"github.com/fractalqb/change"
	"github.com/fractalqb/change/chgv"
)

onchgtype(bool, Bool)
onchgtype(uint8, Uint8)
onchgtype(uint16, Uint16)
onchgtype(uint32, Uint32)
onchgtype(uint64, Uint64)
onchgtype(int8, Int8)
onchgtype(int16, Int16)
onchgtype(int32, Int32)
onchgtype(int64, Int64)
onchgtype(float32, Float32)
onchgtype(float64, Float64)
onchgtype(complex64, Complex64)
onchgtype(complex128, Complex128)
onchgtype(byte, Byte)
onchgtype(rune, Rune)
onchgtype(uint, Uint)
onchgtype(int, Int)
onchgtype(uintptr, UintPtr)
onchgtype(string, String)
