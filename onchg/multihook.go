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

package onchg

import (
	"time"

	"github.com/fractalqb/change"
)

type BoolMultiHook []BoolHook

func (mh BoolMultiHook) Update(src *Bool, ov, nv bool, check bool) (chg change.Flags) {
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

type Uint8MultiHook []Uint8Hook

func (mh Uint8MultiHook) Update(src *Uint8, ov, nv uint8, check bool) (chg change.Flags) {
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

type Uint16MultiHook []Uint16Hook

func (mh Uint16MultiHook) Update(src *Uint16, ov, nv uint16, check bool) (chg change.Flags) {
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

type Uint32MultiHook []Uint32Hook

func (mh Uint32MultiHook) Update(src *Uint32, ov, nv uint32, check bool) (chg change.Flags) {
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

type Uint64MultiHook []Uint64Hook

func (mh Uint64MultiHook) Update(src *Uint64, ov, nv uint64, check bool) (chg change.Flags) {
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

type Int8MultiHook []Int8Hook

func (mh Int8MultiHook) Update(src *Int8, ov, nv int8, check bool) (chg change.Flags) {
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

type Int16MultiHook []Int16Hook

func (mh Int16MultiHook) Update(src *Int16, ov, nv int16, check bool) (chg change.Flags) {
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

type Int32MultiHook []Int32Hook

func (mh Int32MultiHook) Update(src *Int32, ov, nv int32, check bool) (chg change.Flags) {
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

type Int64MultiHook []Int64Hook

func (mh Int64MultiHook) Update(src *Int64, ov, nv int64, check bool) (chg change.Flags) {
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

type Float32MultiHook []Float32Hook

func (mh Float32MultiHook) Update(src *Float32, ov, nv float32, check bool) (chg change.Flags) {
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

type Float64MultiHook []Float64Hook

func (mh Float64MultiHook) Update(src *Float64, ov, nv float64, check bool) (chg change.Flags) {
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

type Complex64MultiHook []Complex64Hook

func (mh Complex64MultiHook) Update(src *Complex64, ov, nv complex64, check bool) (chg change.Flags) {
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

type Complex128MultiHook []Complex128Hook

func (mh Complex128MultiHook) Update(src *Complex128, ov, nv complex128, check bool) (chg change.Flags) {
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

type ByteMultiHook []ByteHook

func (mh ByteMultiHook) Update(src *Byte, ov, nv byte, check bool) (chg change.Flags) {
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

type RuneMultiHook []RuneHook

func (mh RuneMultiHook) Update(src *Rune, ov, nv rune, check bool) (chg change.Flags) {
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

type UintMultiHook []UintHook

func (mh UintMultiHook) Update(src *Uint, ov, nv uint, check bool) (chg change.Flags) {
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

type IntMultiHook []IntHook

func (mh IntMultiHook) Update(src *Int, ov, nv int, check bool) (chg change.Flags) {
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

type UintPtrMultiHook []UintPtrHook

func (mh UintPtrMultiHook) Update(src *UintPtr, ov, nv uintptr, check bool) (chg change.Flags) {
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

type StringMultiHook []StringHook

func (mh StringMultiHook) Update(src *String, ov, nv string, check bool) (chg change.Flags) {
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

type TimeMultiHook []TimeHook

func (mh TimeMultiHook) Update(src *Time, ov, nv time.Time, check bool) (chg change.Flags) {
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

type DurationMultiHook []DurationHook

func (mh DurationMultiHook) Update(src *Duration, ov, nv time.Duration, check bool) (chg change.Flags) {
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
