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

import "fmt"

func ExampleOnChg() {
	user := struct {
		Name   OnChg[string]
		Logins OnChg[int]
	}{
		Name: *NewOnChg[string]("John Doe", FlagHook[string](1).Flag),
		Logins: *NewOnChg[int](0, func(_ *OnChg[int], o, n int, check bool) Flags {
			if !check {
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
