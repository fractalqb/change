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
	"github.com/fractalqb/change/chgv"
)

// time.Time returns chg as change.Flags
func (chg ChgFlag) Time(src *Time, _, _ bool, _ bool) change.Flags {
	return change.Flags(chg)
}

// time.Duration returns chg as change.Flags
func (chg ChgFlag) Duration(src *Duration, _, _ bool, _ bool) change.Flags {
	return change.Flags(chg)
}

// Time tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Time struct {
	v chgv.Time
	h TimeHook
}

func NewTime(init time.Time, hook TimeHook) *Time {
	return &Time{v: chgv.Time(init), h: hook}
}

type TimeHook func(src *Time, ov, nv time.Time, check bool) change.Flags

func (cv *Time) Reset(v time.Time, hook TimeHook) (time.Time, TimeHook) {
	ov, oh := cv.v, cv.h
	(*cv) = Time{v: chgv.Time(v), h: hook}
	return ov.Get(), oh
}

func (cv Time) Get() time.Time { return cv.v.Get() }

func (cv *Time) Set(v time.Time, chg change.Flags) change.Flags {
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

// Duration tracks changes of bool values with a hook and Flags.
// See also onchg package documentation.
type Duration struct {
	v chgv.Duration
	h DurationHook
}

func NewDuration(init time.Duration, hook DurationHook) *Duration {
	return &Duration{v: chgv.Duration(init), h: hook}
}

type DurationHook func(src *Duration, ov, nv time.Duration, check bool) change.Flags

func (cv *Duration) Reset(v time.Duration, hook DurationHook) (time.Duration, DurationHook) {
	ov, oh := cv.v, cv.h
	(*cv) = Duration{v: chgv.Duration(v), h: hook}
	return ov.Get(), oh
}

func (cv Duration) Get() time.Duration { return cv.v.Get() }

func (cv *Duration) Set(v time.Duration, chg change.Flags) change.Flags {
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
