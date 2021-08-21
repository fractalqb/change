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
	"testing"
	"time"

	"github.com/fractalqb/change"
)

var (
	_ change.Time     = (*Time)(nil)
	_ change.Duration = (*Duration)(nil)
)

func TestTime(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		raw := time.Now()
		cv := Time(raw)
		if rs, cs := raw.String(), cv.String(); rs != cs {
			t.Errorf("%s != %s", rs, cs)
		}
	})
	t.Run("zero value", func(t *testing.T) {
		var cv Time
		if !cv.Get().IsZero() {
			t.Error("Int does not initialize to zero value")
		}
	})
	now := time.Now()
	t.Run("init", func(t *testing.T) {
		cv := Time(now)
		if v := cv.Get(); !v.Equal(now) {
			t.Errorf("Int initialization failed. Got %v want %v", v, now)
		}
	})
	t.Run("unchanged set", func(t *testing.T) {
		cv := Time(now)
		chg := cv.Set(now, 1)
		if v := cv.Get(); !v.Equal(now) {
			t.Errorf("Got unexpeted value %v, want %v", v, now)
		}
		if chg != 0 {
			t.Errorf("Unexpected change flags 0x%x, want 0", chg)
		}
	})
	t.Run("changing set", func(t *testing.T) {
		cv := Time(now)
		later := now.Add(time.Minute)
		chg := cv.Set(later, 1)
		if v := cv.Get(); !v.Equal(later) {
			t.Errorf("Got unexpeted value %v, want %v", v, later)
		}
		if chg != 1 {
			t.Errorf("Unexpected change flags 0x%x, want 1", chg)
		}
	})
}

func TestDuration(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		raw := 3 * time.Minute
		cv := Duration(raw)
		if rs, cs := raw.String(), cv.String(); rs != cs {
			t.Errorf("%s != %s", rs, cs)
		}
	})
	t.Run("zero value", func(t *testing.T) {
		var cv Duration
		if cv != 0 {
			t.Error("Int does not initialize to zero value")
		}
	})
	t.Run("init", func(t *testing.T) {
		cv := Duration(time.Second)
		if v := cv.Get(); v != time.Second {
			t.Errorf("Int initialization failed. Got %v want 1s", v)
		}
	})
	t.Run("unchanged set", func(t *testing.T) {
		cv := Duration(time.Second)
		chg := cv.Set(time.Second, 1)
		if v := cv.Get(); v != time.Second {
			t.Errorf("Got unexpeted value %v, want 1s", v)
		}
		if chg != 0 {
			t.Errorf("Unexpected change flags 0x%x, want 0", chg)
		}
	})
	t.Run("changing set", func(t *testing.T) {
		cv := Duration(time.Second)
		chg := cv.Set(2*time.Second, 1)
		if v := cv.Get(); v != 2*time.Second {
			t.Errorf("Got unexpeted value %v, want 2s", v)
		}
		if chg != 1 {
			t.Errorf("Unexpected change flags 0x%x, want 1", chg)
		}
	})
}
