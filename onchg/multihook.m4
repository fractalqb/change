changequote(`<[', `]>')dnl
define(<[mhooktype]>, <[type $2MultiHook []$2Hook

func (mh $2MultiHook) Update(src *$2, ov, nv $1, check bool) (chg change.Flags) {
	if check {
		for _, h := range mh {
			c := h(src, ov, nv, true)
			if c == 0 {
				return 0
			}
			chg |= c
		}
	} else {
		for _, h := range mh {
			chg |= h(src, ov, nv, false)
		}
	}
	return chg
}
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

package onchg

import (
	"time"

	"github.com/fractalqb/change"
)

mhooktype(bool, Bool)
mhooktype(uint8, Uint8)
mhooktype(uint16, Uint16)
mhooktype(uint32, Uint32)
mhooktype(uint64, Uint64)
mhooktype(int8, Int8)
mhooktype(int16, Int16)
mhooktype(int32, Int32)
mhooktype(int64, Int64)
mhooktype(float32, Float32)
mhooktype(float64, Float64)
mhooktype(complex64, Complex64)
mhooktype(complex128, Complex128)
mhooktype(byte, Byte)
mhooktype(rune, Rune)
mhooktype(uint, Uint)
mhooktype(int, Int)
mhooktype(uintptr, UintPtr)
mhooktype(string, String)
dnl
mhooktype(time.Time, Time)
mhooktype(time.Duration, Duration)
