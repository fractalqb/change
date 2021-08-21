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

func TestObserveFields(t *testing.T) {
	type SubPlain struct {
		C Bool
	}
	type SubObl struct {
		ObservableBase
		B Int
	}
	top := struct {
		ObservableBase
		SubPlain
		SubObl
		A String
	}{}
	err := ObserveFields(&top, 0, true, change.AllFlags)
	if err != nil {
		t.Fatal(err)
	}
	var event FieldEvent
	top.ObsRegister(0, "top", &UpdateFunc{func(tag interface{}, e Event) {
		t.Logf("Event: %v", e)
		if tag != "top" {
			t.Errorf("Unexpected tag: '%v'", tag)
		}
		var ok bool
		if event, ok = e.(FieldEvent); !ok {
			t.Errorf("Wrong event type %[1]T: %[1]v", e)
		}
	}})
	chg := top.A.Set("top::A", 0)
	if chg != 4 {
		t.Error("Unxpected change flag for top::A:", chg)
	}
	if event.Field() != "A" {
		t.Error("Event on wrong field:", event)
	}
	chg = top.B.Set(4711, 0)
	if chg != 2 {
		t.Error("Unxpected change flag for top::B:", chg)
	}
	if event.Field() != "SubObl" {
		t.Error("Event on wrong field:", event)
	}
	chg = top.C.Set(true, 0)
	if chg != 1 {
		t.Error("Unxpected change flag for top::C:", chg)
	}
	if event.Field() != "C" {
		t.Error("Event on wrong field:", event)
	}
}
