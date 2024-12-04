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
	"errors"

	"github.com/fractalqb/change"
)

func Scan[T comparable](v change.Able[T], chg change.Flags) *nonNullScanner[T] {
	return &nonNullScanner[T]{v, chg}
}

type nonNullScanner[T comparable] struct {
	V   change.Able[T]
	Chg change.Flags
}

func (scn *nonNullScanner[T]) Scan(src any) error {
	if src == nil {
		return errors.New("scanning illegal null value")
	}
	pv, err := promoteSqlTo[T](src)
	if err != nil {
		return err
	}
	scn.Chg = scn.V.Set(pv, scn.Chg)
	return nil
}

func ScanNull[T comparable](v change.Able[T], null T, chg change.Flags) *nullScanner[T] {
	return &nullScanner[T]{v, null, chg}
}

type nullScanner[T comparable] struct {
	V    change.Able[T]
	Null T
	Chg  change.Flags
}

func (scn *nullScanner[T]) Scan(src any) error {
	if src == nil {
		scn.Chg = scn.V.Set(scn.Null, scn.Chg)
		return nil
	}
	pv, err := promoteSqlTo[T](src)
	if err != nil {
		return err
	}
	scn.Chg = scn.V.Set(pv, scn.Chg)
	return nil
}
