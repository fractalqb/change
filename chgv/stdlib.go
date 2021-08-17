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

import (
	"time"

	"github.com/fractalqb/change"
)

// Time tracks changes of time.Time values with caller-provided Flags.
// See also chgv package documentation.
type Time time.Time

func (cv Time) Get() time.Time { return time.Time(cv) }

func (cv *Time) Set(v time.Time, chg change.Flags) change.Flags {
	if time.Time(*cv).Equal(v) {
		return 0
	}
	*cv = Time(v)
	return chg
}

func (cv Time) String() string { return time.Time(cv).String() }

// Duration tracks changes of time.Duration values with
// caller-provided Flags.  See also chgv package documentation.
type Duration time.Duration

func (cv Duration) Get() time.Duration { return time.Duration(cv) }

func (cv *Duration) Set(v time.Duration, chg change.Flags) change.Flags {
	if time.Duration(*cv) == v {
		return 0
	}
	*cv = Duration(v)
	return chg
}

func (cv Duration) String() string { return time.Duration(cv).String() }
