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

package chgv

import (
	"testing"

	"github.com/fractalqb/change"
)

func BenchmarkIntReference(b *testing.B) {
	var val int
	for i := 0; i < b.N; i++ {
		tmp := val
		tmp += i
		val = tmp
	}
}

func BenchmarkInt_noDetect(b *testing.B) {
	var val Int
	for i := 0; i < b.N; i++ {
		tmp := val.Get()
		tmp += i
		val.Set(tmp, 0)
	}
}

func BenchmarkInt_withDetect(b *testing.B) {
	var val Int
	var chg change.Flags
	count := 0
	for i := 0; i < b.N; i++ {
		tmp := val.Get()
		tmp += i
		c := val.Set(tmp, 0)
		if c != 0 {
			count++
		}
		chg |= c
	}
}
