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

import "github.com/fractalqb/change"

type ValScanner[T comparable] struct {
	V *change.Val[T]
	Chg change.Flags
	Null T
}

func (v ValScanner[T]) Scan(src interface{}) error {
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
