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

	"github.com/fractalqb/change"
)

type OnScanner[T comparable] struct {
	V    *change.On[T]
	Chg  change.Flags
	Null T
}

func (v *OnScanner[T]) Scan(src any) error {
	if src == nil {
		v.Chg = v.V.Set(v.Null, v.Chg)
		return nil
	}
	pv, err := promoteSqlTo[T](src)
	if err != nil {
		return err
	}
	v.Chg = v.V.Set(pv, v.Chg)
	return nil
}

type OnScanF32 struct {
	V   *change.On[float32]
	Chg change.Flags
}

func (v *OnScanF32) Scan(src any) error {
	if src == nil {
		v.Chg = v.V.Set(nan32, v.Chg)
		return nil
	}
	switch src := src.(type) {
	case float64:
		v.Chg = v.V.Set(float32(src), v.Chg)
	case float32:
		v.Chg = v.V.Set(src, v.Chg)
	default:
		return fmt.Errorf("cannot promote SQL %T to float32", src)
	}
	return nil
}

type OnScanF64 struct {
	V   *change.On[float64]
	Chg change.Flags
}

func (v *OnScanF64) Scan(src any) error {
	if src == nil {
		v.Chg = v.V.Set(math.NaN(), v.Chg)
		return nil
	}
	switch src := src.(type) {
	case float64:
		v.Chg = v.V.Set(src, v.Chg)
	case float32:
		v.Chg = v.V.Set(float64(src), v.Chg)
	default:
		return fmt.Errorf("cannot promote SQL %T to float64", src)
	}
	return nil
}
