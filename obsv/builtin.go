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
	"fmt"

	"github.com/fractalqb/change"
	"github.com/fractalqb/change/chgv"
)

type ValueChange interface {
	Event
	OldValue() interface{}
	NewValue() interface{}
}

type Bool struct {
	v chgv.Bool
	ObservableBase
}

func NewBool(init bool, defaultTag interface{}, defaultChg change.Flags) *Bool {
	res := &Bool{v: chgv.Bool(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Bool) Get() bool { return cv.v.Get() }

func (cv *Bool) Set(v bool, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := BoolChg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type BoolChg struct {
	eventBase
	ov, nv bool
}

func (ic BoolChg) String() string {
	return fmt.Sprintf("Chg bool [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic BoolChg) OldValue() interface{} { return ic.ov }
func (ic BoolChg) NewValue() interface{} { return ic.nv }

type Uint8 struct {
	v chgv.Uint8
	ObservableBase
}

func NewUint8(init uint8, defaultTag interface{}, defaultChg change.Flags) *Uint8 {
	res := &Uint8{v: chgv.Uint8(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Uint8) Get() uint8 { return cv.v.Get() }

func (cv *Uint8) Set(v uint8, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Uint8Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Uint8Chg struct {
	eventBase
	ov, nv uint8
}

func (ic Uint8Chg) String() string {
	return fmt.Sprintf("Chg uint8 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Uint8Chg) OldValue() interface{} { return ic.ov }
func (ic Uint8Chg) NewValue() interface{} { return ic.nv }

type Uint16 struct {
	v chgv.Uint16
	ObservableBase
}

func NewUint16(init uint16, defaultTag interface{}, defaultChg change.Flags) *Uint16 {
	res := &Uint16{v: chgv.Uint16(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Uint16) Get() uint16 { return cv.v.Get() }

func (cv *Uint16) Set(v uint16, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Uint16Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Uint16Chg struct {
	eventBase
	ov, nv uint16
}

func (ic Uint16Chg) String() string {
	return fmt.Sprintf("Chg uint16 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Uint16Chg) OldValue() interface{} { return ic.ov }
func (ic Uint16Chg) NewValue() interface{} { return ic.nv }

type Uint32 struct {
	v chgv.Uint32
	ObservableBase
}

func NewUint32(init uint32, defaultTag interface{}, defaultChg change.Flags) *Uint32 {
	res := &Uint32{v: chgv.Uint32(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Uint32) Get() uint32 { return cv.v.Get() }

func (cv *Uint32) Set(v uint32, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Uint32Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Uint32Chg struct {
	eventBase
	ov, nv uint32
}

func (ic Uint32Chg) String() string {
	return fmt.Sprintf("Chg uint32 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Uint32Chg) OldValue() interface{} { return ic.ov }
func (ic Uint32Chg) NewValue() interface{} { return ic.nv }

type Uint64 struct {
	v chgv.Uint64
	ObservableBase
}

func NewUint64(init uint64, defaultTag interface{}, defaultChg change.Flags) *Uint64 {
	res := &Uint64{v: chgv.Uint64(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Uint64) Get() uint64 { return cv.v.Get() }

func (cv *Uint64) Set(v uint64, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Uint64Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Uint64Chg struct {
	eventBase
	ov, nv uint64
}

func (ic Uint64Chg) String() string {
	return fmt.Sprintf("Chg uint64 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Uint64Chg) OldValue() interface{} { return ic.ov }
func (ic Uint64Chg) NewValue() interface{} { return ic.nv }

type Int8 struct {
	v chgv.Int8
	ObservableBase
}

func NewInt8(init int8, defaultTag interface{}, defaultChg change.Flags) *Int8 {
	res := &Int8{v: chgv.Int8(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Int8) Get() int8 { return cv.v.Get() }

func (cv *Int8) Set(v int8, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Int8Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Int8Chg struct {
	eventBase
	ov, nv int8
}

func (ic Int8Chg) String() string {
	return fmt.Sprintf("Chg int8 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Int8Chg) OldValue() interface{} { return ic.ov }
func (ic Int8Chg) NewValue() interface{} { return ic.nv }

type Int16 struct {
	v chgv.Int16
	ObservableBase
}

func NewInt16(init int16, defaultTag interface{}, defaultChg change.Flags) *Int16 {
	res := &Int16{v: chgv.Int16(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Int16) Get() int16 { return cv.v.Get() }

func (cv *Int16) Set(v int16, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Int16Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Int16Chg struct {
	eventBase
	ov, nv int16
}

func (ic Int16Chg) String() string {
	return fmt.Sprintf("Chg int16 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Int16Chg) OldValue() interface{} { return ic.ov }
func (ic Int16Chg) NewValue() interface{} { return ic.nv }

type Int32 struct {
	v chgv.Int32
	ObservableBase
}

func NewInt32(init int32, defaultTag interface{}, defaultChg change.Flags) *Int32 {
	res := &Int32{v: chgv.Int32(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Int32) Get() int32 { return cv.v.Get() }

func (cv *Int32) Set(v int32, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Int32Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Int32Chg struct {
	eventBase
	ov, nv int32
}

func (ic Int32Chg) String() string {
	return fmt.Sprintf("Chg int32 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Int32Chg) OldValue() interface{} { return ic.ov }
func (ic Int32Chg) NewValue() interface{} { return ic.nv }

type Int64 struct {
	v chgv.Int64
	ObservableBase
}

func NewInt64(init int64, defaultTag interface{}, defaultChg change.Flags) *Int64 {
	res := &Int64{v: chgv.Int64(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Int64) Get() int64 { return cv.v.Get() }

func (cv *Int64) Set(v int64, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Int64Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Int64Chg struct {
	eventBase
	ov, nv int64
}

func (ic Int64Chg) String() string {
	return fmt.Sprintf("Chg int64 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Int64Chg) OldValue() interface{} { return ic.ov }
func (ic Int64Chg) NewValue() interface{} { return ic.nv }

type Float32 struct {
	v chgv.Float32
	ObservableBase
}

func NewFloat32(init float32, defaultTag interface{}, defaultChg change.Flags) *Float32 {
	res := &Float32{v: chgv.Float32(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Float32) Get() float32 { return cv.v.Get() }

func (cv *Float32) Set(v float32, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Float32Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Float32Chg struct {
	eventBase
	ov, nv float32
}

func (ic Float32Chg) String() string {
	return fmt.Sprintf("Chg float32 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Float32Chg) OldValue() interface{} { return ic.ov }
func (ic Float32Chg) NewValue() interface{} { return ic.nv }

type Float64 struct {
	v chgv.Float64
	ObservableBase
}

func NewFloat64(init float64, defaultTag interface{}, defaultChg change.Flags) *Float64 {
	res := &Float64{v: chgv.Float64(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Float64) Get() float64 { return cv.v.Get() }

func (cv *Float64) Set(v float64, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Float64Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Float64Chg struct {
	eventBase
	ov, nv float64
}

func (ic Float64Chg) String() string {
	return fmt.Sprintf("Chg float64 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Float64Chg) OldValue() interface{} { return ic.ov }
func (ic Float64Chg) NewValue() interface{} { return ic.nv }

type Complex64 struct {
	v chgv.Complex64
	ObservableBase
}

func NewComplex64(init complex64, defaultTag interface{}, defaultChg change.Flags) *Complex64 {
	res := &Complex64{v: chgv.Complex64(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Complex64) Get() complex64 { return cv.v.Get() }

func (cv *Complex64) Set(v complex64, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Complex64Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Complex64Chg struct {
	eventBase
	ov, nv complex64
}

func (ic Complex64Chg) String() string {
	return fmt.Sprintf("Chg complex64 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Complex64Chg) OldValue() interface{} { return ic.ov }
func (ic Complex64Chg) NewValue() interface{} { return ic.nv }

type Complex128 struct {
	v chgv.Complex128
	ObservableBase
}

func NewComplex128(init complex128, defaultTag interface{}, defaultChg change.Flags) *Complex128 {
	res := &Complex128{v: chgv.Complex128(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Complex128) Get() complex128 { return cv.v.Get() }

func (cv *Complex128) Set(v complex128, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := Complex128Chg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type Complex128Chg struct {
	eventBase
	ov, nv complex128
}

func (ic Complex128Chg) String() string {
	return fmt.Sprintf("Chg complex128 [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic Complex128Chg) OldValue() interface{} { return ic.ov }
func (ic Complex128Chg) NewValue() interface{} { return ic.nv }

type Byte struct {
	v chgv.Byte
	ObservableBase
}

func NewByte(init byte, defaultTag interface{}, defaultChg change.Flags) *Byte {
	res := &Byte{v: chgv.Byte(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Byte) Get() byte { return cv.v.Get() }

func (cv *Byte) Set(v byte, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := ByteChg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type ByteChg struct {
	eventBase
	ov, nv byte
}

func (ic ByteChg) String() string {
	return fmt.Sprintf("Chg byte [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic ByteChg) OldValue() interface{} { return ic.ov }
func (ic ByteChg) NewValue() interface{} { return ic.nv }

type Rune struct {
	v chgv.Rune
	ObservableBase
}

func NewRune(init rune, defaultTag interface{}, defaultChg change.Flags) *Rune {
	res := &Rune{v: chgv.Rune(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Rune) Get() rune { return cv.v.Get() }

func (cv *Rune) Set(v rune, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := RuneChg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type RuneChg struct {
	eventBase
	ov, nv rune
}

func (ic RuneChg) String() string {
	return fmt.Sprintf("Chg rune [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic RuneChg) OldValue() interface{} { return ic.ov }
func (ic RuneChg) NewValue() interface{} { return ic.nv }

type Uint struct {
	v chgv.Uint
	ObservableBase
}

func NewUint(init uint, defaultTag interface{}, defaultChg change.Flags) *Uint {
	res := &Uint{v: chgv.Uint(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Uint) Get() uint { return cv.v.Get() }

func (cv *Uint) Set(v uint, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := UintChg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type UintChg struct {
	eventBase
	ov, nv uint
}

func (ic UintChg) String() string {
	return fmt.Sprintf("Chg uint [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic UintChg) OldValue() interface{} { return ic.ov }
func (ic UintChg) NewValue() interface{} { return ic.nv }

type Int struct {
	v chgv.Int
	ObservableBase
}

func NewInt(init int, defaultTag interface{}, defaultChg change.Flags) *Int {
	res := &Int{v: chgv.Int(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv Int) Get() int { return cv.v.Get() }

func (cv *Int) Set(v int, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := IntChg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type IntChg struct {
	eventBase
	ov, nv int
}

func (ic IntChg) String() string {
	return fmt.Sprintf("Chg int [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic IntChg) OldValue() interface{} { return ic.ov }
func (ic IntChg) NewValue() interface{} { return ic.nv }

type UintPtr struct {
	v chgv.UintPtr
	ObservableBase
}

func NewUintPtr(init uintptr, defaultTag interface{}, defaultChg change.Flags) *UintPtr {
	res := &UintPtr{v: chgv.UintPtr(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv UintPtr) Get() uintptr { return cv.v.Get() }

func (cv *UintPtr) Set(v uintptr, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := UintPtrChg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type UintPtrChg struct {
	eventBase
	ov, nv uintptr
}

func (ic UintPtrChg) String() string {
	return fmt.Sprintf("Chg uintptr [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic UintPtrChg) OldValue() interface{} { return ic.ov }
func (ic UintPtrChg) NewValue() interface{} { return ic.nv }

type String struct {
	v chgv.String
	ObservableBase
}

func NewString(init string, defaultTag interface{}, defaultChg change.Flags) *String {
	res := &String{v: chgv.String(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

func (cv String) Get() string { return cv.v.Get() }

func (cv *String) Set(v string, chg change.Flags) change.Flags {
	if cv.stub == nil {
		return cv.v.Set(v, chg)
	}
	cv.stub.veto = nil
	e := StringChg{
		eventBase: eventBase{
			src: cv,
			chg: cv.stub.change(chg),
		},
		ov: cv.v.Get(),
		nv: v,
	}
	if c, veto := cv.stub.check(e); veto != nil {
		cv.stub.veto = veto
		return 0
	} else if chg == 0 {
		chg = c
	}
	if chg = cv.v.Set(v, chg); chg == 0 {
		return 0
	}
	e.chg = chg
	cv.stub.notify(e)
	return chg
}

type StringChg struct {
	eventBase
	ov, nv string
}

func (ic StringChg) String() string {
	return fmt.Sprintf("Chg string [%v] -> [%v] flags=%x",
		ic.ov, ic.nv,
		ic.chg)
}

func (ic StringChg) OldValue() interface{} { return ic.ov }
func (ic StringChg) NewValue() interface{} { return ic.nv }
