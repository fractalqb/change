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

type fieldUpdate struct {
	object Observable
	field  string
}

func (f fieldUpdate) Check(_ interface{}, _ Event) (change.Flags, error) {
	return 0, nil
}

func (f fieldUpdate) Update(tag interface{}, e Event) {
	fe := FieldEvent{Event: e, fieldTag: tag, field: f.field}
	f.object.ObsEach(func(tag interface{}, o Observer) {
		o.Update(tag, fe)
	})
}

// ObserveFields registers the struct obj as an Observer to all its Observable
// fields so that updates of the fields are forwarded to obj's Observers as
// FieldEvents.
//
// Field Tags
//
// TODO doc `chgv:",tag=some-tag,flag=0x10"`
func ObserveFields(obj Observable, prio int, reflectedFieldTag bool, flags change.Flags) error {
	val := reflect.Indirect(reflect.ValueOf(obj))
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("Cannot observe fields of non-struct type %T", obj)
	}
	type oblfield struct {
		sf reflect.StructField
		o  Observable
	}
	var oblfields []oblfield
	eachOblField(val, func(sf reflect.StructField, o Observable) {
		oblfields = append(oblfields, oblfield{sf, o})
	})
	stubs := make([]stub, len(oblfields))
	setFlags := flags != 0
	for fi, of := range oblfields {
		if of.sf.Anonymous && of.sf.Type.Kind() == reflect.Struct {
			err := ObserveFields(of.o, prio, reflectedFieldTag, flags)
			if err != nil {
				return err
			}
		}
		field := of.sf.Name
		var (
			defaultTag  interface{}
			defaultFlag change.Flags
		)
		if tag, ok := of.sf.Tag.Lookup("obsv"); ok {
			if tag == "-" {
				continue
			}
			var err error
			field, defaultTag, defaultFlag, err = paseFieldTag(field, tag)
			if err != nil {
				return fmt.Errorf("Tag on field '%s' in %T: %w",
					of.sf.Name,
					obj,
					err)
			}
			if flags&defaultFlag != defaultFlag {
				return fmt.Errorf("Flag 0x%x not available for field '%s' in %T",
					defaultFlag,
					of.sf.Name,
					obj)
			}
		}
		var regTag interface{}
		if reflectedFieldTag {
			regTag = of.sf
		}
		of.o.register(prio, regTag, fieldUpdate{
			object: obj,
			field:  field,
		}, &stubs[fi])
		if setFlags && defaultFlag == 0 {
			if flags == 0 {
				return fmt.Errorf("No flags left for field '%s'", of.sf.Name)
			}
			defaultFlag, flags = pickFlag(flags)
		} else {
			flags &= ^defaultFlag
		}
		of.o.SetObsDefaults(defaultTag, defaultFlag)
	}
	return nil
}

var oblBaseType = reflect.TypeOf(ObservableBase{})

func eachOblField(obj reflect.Value, do func(reflect.StructField, Observable)) {
	typ := obj.Type()
	for i := 0; i < obj.NumField(); i++ {
		fv := obj.Field(i)
		ft := fv.Type()
		sf := typ.Field(i)
		if ft == oblBaseType || !sf.IsExported() {
			continue
		}
		if fobs := isObservable(fv); fobs != nil {
			do(sf, fobs)
		} else if sf.Anonymous && ft.Kind() == reflect.Struct {
			eachOblField(fv, do)
		}
	}
}

func isObservable(v reflect.Value) Observable {
	if o, ok := v.Interface().(Observable); ok {
		return o
	}
	if v.Kind() == reflect.Ptr {
		return nil
	}
	v = v.Addr()
	if o, ok := v.Interface().(Observable); ok {
		return o
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
