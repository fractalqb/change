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
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/fractalqb/change"
)

type FieldEvent struct {
	field    string
	fieldTag interface{}
	Event
}

func (fe FieldEvent) Field() string { return fe.field }

func (fe FieldEvent) FieldTag() interface{} { return fe.fieldTag }

func (fe FieldEvent) String() string {
	return fmt.Sprintf("Field '%s': tag=%+v event=%s",
		fe.field,
		fe.fieldTag,
		fe.Event)
}

type FieldUpdate struct {
	Object Observable
	Field  string
}

func (f FieldUpdate) Check(_ interface{}, _ Event) (change.Flags, error) {
	return 0, nil
}

func (f FieldUpdate) Update(tag interface{}, e Event) {
	fe := FieldEvent{Event: e, fieldTag: tag, field: f.Field}
	f.Object.ObsEach(func(tag interface{}, o Observer) {
		o.Update(tag, fe)
	})
}

// TODO doc `chgv:",tag=some-tag,flag=0x10"`
func ObserveFields(obj Observable, prio int, reflectedFieldTag bool, flags change.Flags) error {
	val := reflect.Indirect(reflect.ValueOf(obj))
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("Cannot observe fields of non-struct type %T", obj)
	}
	setFlags := flags != 0
	typ := val.Type()
	obsBaseType := reflect.TypeOf(ObservableBase{})
	for i := 0; i < val.NumField(); i++ {
		fv := val.Field(i)
		if fv.Type() == obsBaseType {
			continue
		}
		if fv.Kind() != reflect.Ptr {
			fv = fv.Addr()
		}
		if fobs, ok := fv.Interface().(Observable); ok {
			sf := typ.Field(i)
			field := sf.Name
			var (
				defaultTag  interface{}
				defaultFlag change.Flags
			)
			if tag, ok := sf.Tag.Lookup("obsv"); ok {
				if tag == "-" {
					continue
				}
				var err error
				field, defaultTag, defaultFlag, err = paseFieldTag(field, tag)
				if err != nil {
					return fmt.Errorf("Tag on field '%s' in %T: %w",
						sf.Name,
						obj,
						err)
				}
				if flags&defaultFlag != defaultFlag {
					return fmt.Errorf("Flag 0x%x not available for field '%s' in %T",
						defaultFlag,
						sf.Name,
						obj)
				}
			}
			var regTag interface{}
			if reflectedFieldTag {
				regTag = sf
			}
			fobs.ObsRegister(prio, regTag, FieldUpdate{
				Object: obj,
				Field:  field,
			})
			if setFlags && defaultFlag == 0 {
				if flags == 0 {
					return fmt.Errorf("No flags left for field %d '%s'",
						i,
						sf.Name)
				}
				defaultFlag, flags = pickFlag(flags)
			} else {
				flags &= ^defaultFlag
			}
			fobs.SetObsDefaults(defaultTag, defaultFlag)
		}
	}
	return nil
}

func paseFieldTag(f, ftag string) (field, tag string, flag change.Flags, err error) {
	if ftag == "" {
		return "", "", 0, errors.New("Empty tag value")
	}
	fields := strings.Split(ftag, ",")
	if fields[0] != "" {
		field = fields[0]
	} else {
		field = f
	}
	for _, fval := range fields[1:] {
		switch {
		case strings.HasPrefix(fval, "tag="):
			tag = fval[4:]
		case strings.HasPrefix(fval, "flag="):
			f, err := strconv.ParseUint(fval[5:], 0, 64)
			if err != nil {
				return field, tag, 0, fmt.Errorf("flag: %v", err)
			}
			if f == 0 {
				return field, tag, 0, fmt.Errorf("Cannot set 0 as flag")
			}
			flag = change.Flags(f)
		}
	}
	return field, tag, flag, nil
}

func pickFlag(from change.Flags) (flag, rest change.Flags) {
	for flag = change.Flags(1); flag != 0; flag <<= 1 {
		if from&flag == flag {
			return flag, from & ^flag
		}
	}
	return 0, from
}
