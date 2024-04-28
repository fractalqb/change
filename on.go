// Copyright 2021 Marcus Perlick
// This file is part of Go module github.com/fractalqb/change
//
// change is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// change is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with change.  If not, see <http://www.gnu.org/licenses/>.

package change

// On implements [Changeable] with a hook that is called on value changes.  When
// [On.Set] is called, first the hook is called before a change is made with
// check=true. With check==true the hook decides if it blocks the set operation.
// To block the set operation, the hook returns 0. Otherwise the hook provides
// the change flags for the Set method. The flags from the hook are overridden
// if non-zero flags are passed to the value's Set method. If the hook is nil,
// On behave exactly like [Val].
type On[T comparable] struct {
	v Val[T]
	h HookFunc[T]
}

var _ Changeable[int] = (*On[int])(nil)

// HookFunc functions can be hooked into [On] values. They get passed the On
// object src for which the Set method was called, the old value odlval and the
// to be set value newval and the check flag. See also the description of On.
type HookFunc[T comparable] func(src *On[T], oldval, newval T, check bool) Flags

// FlagHook's Flag method always returns the same Flags value.
type FlagHook[_ comparable] Flags

// Flag is a HookFunc for On values.
func (h FlagHook[T]) Func(src *On[T], oldval, newval T, check bool) Flags {
	return Flags(h)
}

func NewOn[T comparable](init T, hook HookFunc[T]) On[T] {
	return On[T]{v: Val[T]{init}, h: hook}
}

func (cv *On[T]) Reset(v T, hook HookFunc[T]) (T, HookFunc[T]) {
	ov, oh := cv.v.Get(), cv.h
	(*cv) = On[T]{v: Val[T]{v}, h: hook}
	return ov, oh
}

func (cv On[T]) Get() T { return cv.v.Get() }

func (cv *On[T]) Set(v T, chg Flags) Flags {
	if cv.h == nil {
		return cv.v.Set(v, chg)
	}
	return cv.slowSet(v, chg)
}

func (cv *On[T]) slowSet(v T, chg Flags) Flags {
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
