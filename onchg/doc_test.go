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
	"fmt"

	"github.com/fractalqb/change"
)

func Example() {
	user := struct {
		Name   String
		Logins Int
	}{
		Name: *NewString("John Doe", ChgFlag(1).String),
		Logins: *NewInt(0, func(i *Int, o, n int, pre bool) change.Flags {
			if !pre {
				fmt.Printf("changed logins from %d to %d\n", o, n)
			}
			return 2
		}),
	}

	chg := user.Name.Set("John Doe", 0)
	chg |= user.Logins.Set(1, 0)

	fmt.Printf("Changes: 0b%b\n", chg)
	if chg&1 == 0 {
		fmt.Println("Name did not change")
	} else {
		fmt.Println("Name changed")
	}
	// Output:
	// changed logins from 0 to 1
	// Changes: 0b10
	// Name did not change
}
