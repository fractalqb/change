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
	"testing"

	"github.com/fractalqb/change"
)

func BenchmarkInt_noDetect(b *testing.B) {
	var val Int
	for i := 0; i < b.N; i++ {
		tmp := val.Get()
		tmp += i
		val.Set(tmp, 0)
	}
}

func BenchmarkInt_withDetect(b *testing.B) {
	val := NewInt(0, b.Name(), 1)
	var chg change.Flags
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp := val.Get()
		tmp += i
		chg |= val.Set(tmp, 0)
	}
}

func BenchmarkInt_withObserver(b *testing.B) {
	val := NewInt(0, b.Name(), 1)
	var chg change.Flags
	count := 0
	val.ObsRegister(0, nil, &UpdateFunc{func(tag interface{}, e Event) {
		count++
	}})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tmp := val.Get()
		tmp += i
		chg |= val.Set(tmp, 0)
	}
}
