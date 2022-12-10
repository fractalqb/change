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

package chgsql

import (
	"fmt"
	"math"
	"reflect"

	"github.com/fractalqb/change"
)

func Nullable[T comparable](v change.Changeable[T], null T) any {
	db := v.Get()
	if db == null {
		return nil
	}
	return db
}

func promoteSqlTo[T any](from any) (p T, err error) {
	if p, ok := from.(T); ok {
		return p, nil
	}
	fromVal := reflect.ValueOf(from)
	pType := reflect.TypeOf(p)
	if !fromVal.CanConvert(pType) {
		return p, fmt.Errorf("cannot promote SQL %T to %T", from, p)
	}
	p = fromVal.Convert(pType).Interface().(T)
	return p, nil
}

func NullableF32(v change.Changeable[float32]) any {
	f := v.Get()
	if math.IsNaN(float64(f)) {
		return nil
	}
	return f
}

func NullableF64(v change.Changeable[float64]) any {
	f := v.Get()
	if math.IsNaN(f) {
		return nil
	}
	return f
}

var nan32 = float32(math.NaN())
