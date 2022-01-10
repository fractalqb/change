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
	"fmt"

	"github.com/fractalqb/change"
	"github.com/fractalqb/change/chgv"
)

// Bool is an Observable bool value. Bool implements change.Bool.
type Bool struct {
	v chgv.Bool
	ObservableBase
}

// NewBool creates a change-detectable Bool and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewBool(init bool, defaultTag interface{}, defaultChg change.Flags) *Bool {
	res := &Bool{v: chgv.Bool(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current bool value.
func (cv Bool) Get() bool { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// BoolChg is the value change Event for Bool.
type BoolChg struct {
	eventBase
	ov, nv bool
}

func (ce BoolChg) String() string {
	return fmt.Sprintf("Chg bool [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce BoolChg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce BoolChg) NewValue() interface{} { return ce.nv }

// Uint8 is an Observable uint8 value. Uint8 implements change.Uint8.
type Uint8 struct {
	v chgv.Uint8
	ObservableBase
}

// NewUint8 creates a change-detectable Uint8 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewUint8(init uint8, defaultTag interface{}, defaultChg change.Flags) *Uint8 {
	res := &Uint8{v: chgv.Uint8(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current uint8 value.
func (cv Uint8) Get() uint8 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Uint8Chg is the value change Event for Uint8.
type Uint8Chg struct {
	eventBase
	ov, nv uint8
}

func (ce Uint8Chg) String() string {
	return fmt.Sprintf("Chg uint8 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Uint8Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Uint8Chg) NewValue() interface{} { return ce.nv }

// Uint16 is an Observable uint16 value. Uint16 implements change.Uint16.
type Uint16 struct {
	v chgv.Uint16
	ObservableBase
}

// NewUint16 creates a change-detectable Uint16 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewUint16(init uint16, defaultTag interface{}, defaultChg change.Flags) *Uint16 {
	res := &Uint16{v: chgv.Uint16(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current uint16 value.
func (cv Uint16) Get() uint16 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Uint16Chg is the value change Event for Uint16.
type Uint16Chg struct {
	eventBase
	ov, nv uint16
}

func (ce Uint16Chg) String() string {
	return fmt.Sprintf("Chg uint16 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Uint16Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Uint16Chg) NewValue() interface{} { return ce.nv }

// Uint32 is an Observable uint32 value. Uint32 implements change.Uint32.
type Uint32 struct {
	v chgv.Uint32
	ObservableBase
}

// NewUint32 creates a change-detectable Uint32 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewUint32(init uint32, defaultTag interface{}, defaultChg change.Flags) *Uint32 {
	res := &Uint32{v: chgv.Uint32(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current uint32 value.
func (cv Uint32) Get() uint32 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Uint32Chg is the value change Event for Uint32.
type Uint32Chg struct {
	eventBase
	ov, nv uint32
}

func (ce Uint32Chg) String() string {
	return fmt.Sprintf("Chg uint32 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Uint32Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Uint32Chg) NewValue() interface{} { return ce.nv }

// Uint64 is an Observable uint64 value. Uint64 implements change.Uint64.
type Uint64 struct {
	v chgv.Uint64
	ObservableBase
}

// NewUint64 creates a change-detectable Uint64 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewUint64(init uint64, defaultTag interface{}, defaultChg change.Flags) *Uint64 {
	res := &Uint64{v: chgv.Uint64(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current uint64 value.
func (cv Uint64) Get() uint64 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Uint64Chg is the value change Event for Uint64.
type Uint64Chg struct {
	eventBase
	ov, nv uint64
}

func (ce Uint64Chg) String() string {
	return fmt.Sprintf("Chg uint64 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Uint64Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Uint64Chg) NewValue() interface{} { return ce.nv }

// Int8 is an Observable int8 value. Int8 implements change.Int8.
type Int8 struct {
	v chgv.Int8
	ObservableBase
}

// NewInt8 creates a change-detectable Int8 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewInt8(init int8, defaultTag interface{}, defaultChg change.Flags) *Int8 {
	res := &Int8{v: chgv.Int8(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current int8 value.
func (cv Int8) Get() int8 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Int8Chg is the value change Event for Int8.
type Int8Chg struct {
	eventBase
	ov, nv int8
}

func (ce Int8Chg) String() string {
	return fmt.Sprintf("Chg int8 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Int8Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Int8Chg) NewValue() interface{} { return ce.nv }

// Int16 is an Observable int16 value. Int16 implements change.Int16.
type Int16 struct {
	v chgv.Int16
	ObservableBase
}

// NewInt16 creates a change-detectable Int16 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewInt16(init int16, defaultTag interface{}, defaultChg change.Flags) *Int16 {
	res := &Int16{v: chgv.Int16(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current int16 value.
func (cv Int16) Get() int16 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Int16Chg is the value change Event for Int16.
type Int16Chg struct {
	eventBase
	ov, nv int16
}

func (ce Int16Chg) String() string {
	return fmt.Sprintf("Chg int16 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Int16Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Int16Chg) NewValue() interface{} { return ce.nv }

// Int32 is an Observable int32 value. Int32 implements change.Int32.
type Int32 struct {
	v chgv.Int32
	ObservableBase
}

// NewInt32 creates a change-detectable Int32 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewInt32(init int32, defaultTag interface{}, defaultChg change.Flags) *Int32 {
	res := &Int32{v: chgv.Int32(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current int32 value.
func (cv Int32) Get() int32 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Int32Chg is the value change Event for Int32.
type Int32Chg struct {
	eventBase
	ov, nv int32
}

func (ce Int32Chg) String() string {
	return fmt.Sprintf("Chg int32 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Int32Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Int32Chg) NewValue() interface{} { return ce.nv }

// Int64 is an Observable int64 value. Int64 implements change.Int64.
type Int64 struct {
	v chgv.Int64
	ObservableBase
}

// NewInt64 creates a change-detectable Int64 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewInt64(init int64, defaultTag interface{}, defaultChg change.Flags) *Int64 {
	res := &Int64{v: chgv.Int64(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current int64 value.
func (cv Int64) Get() int64 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Int64Chg is the value change Event for Int64.
type Int64Chg struct {
	eventBase
	ov, nv int64
}

func (ce Int64Chg) String() string {
	return fmt.Sprintf("Chg int64 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Int64Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Int64Chg) NewValue() interface{} { return ce.nv }

// Float32 is an Observable float32 value. Float32 implements change.Float32.
type Float32 struct {
	v chgv.Float32
	ObservableBase
}

// NewFloat32 creates a change-detectable Float32 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewFloat32(init float32, defaultTag interface{}, defaultChg change.Flags) *Float32 {
	res := &Float32{v: chgv.Float32(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current float32 value.
func (cv Float32) Get() float32 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Float32Chg is the value change Event for Float32.
type Float32Chg struct {
	eventBase
	ov, nv float32
}

func (ce Float32Chg) String() string {
	return fmt.Sprintf("Chg float32 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Float32Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Float32Chg) NewValue() interface{} { return ce.nv }

// Float64 is an Observable float64 value. Float64 implements change.Float64.
type Float64 struct {
	v chgv.Float64
	ObservableBase
}

// NewFloat64 creates a change-detectable Float64 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewFloat64(init float64, defaultTag interface{}, defaultChg change.Flags) *Float64 {
	res := &Float64{v: chgv.Float64(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current float64 value.
func (cv Float64) Get() float64 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Float64Chg is the value change Event for Float64.
type Float64Chg struct {
	eventBase
	ov, nv float64
}

func (ce Float64Chg) String() string {
	return fmt.Sprintf("Chg float64 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Float64Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Float64Chg) NewValue() interface{} { return ce.nv }

// Complex64 is an Observable complex64 value. Complex64 implements change.Complex64.
type Complex64 struct {
	v chgv.Complex64
	ObservableBase
}

// NewComplex64 creates a change-detectable Complex64 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewComplex64(init complex64, defaultTag interface{}, defaultChg change.Flags) *Complex64 {
	res := &Complex64{v: chgv.Complex64(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current complex64 value.
func (cv Complex64) Get() complex64 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Complex64Chg is the value change Event for Complex64.
type Complex64Chg struct {
	eventBase
	ov, nv complex64
}

func (ce Complex64Chg) String() string {
	return fmt.Sprintf("Chg complex64 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Complex64Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Complex64Chg) NewValue() interface{} { return ce.nv }

// Complex128 is an Observable complex128 value. Complex128 implements change.Complex128.
type Complex128 struct {
	v chgv.Complex128
	ObservableBase
}

// NewComplex128 creates a change-detectable Complex128 and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewComplex128(init complex128, defaultTag interface{}, defaultChg change.Flags) *Complex128 {
	res := &Complex128{v: chgv.Complex128(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current complex128 value.
func (cv Complex128) Get() complex128 { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// Complex128Chg is the value change Event for Complex128.
type Complex128Chg struct {
	eventBase
	ov, nv complex128
}

func (ce Complex128Chg) String() string {
	return fmt.Sprintf("Chg complex128 [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce Complex128Chg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce Complex128Chg) NewValue() interface{} { return ce.nv }

// Byte is an Observable byte value. Byte implements change.Byte.
type Byte struct {
	v chgv.Byte
	ObservableBase
}

// NewByte creates a change-detectable Byte and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewByte(init byte, defaultTag interface{}, defaultChg change.Flags) *Byte {
	res := &Byte{v: chgv.Byte(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current byte value.
func (cv Byte) Get() byte { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// ByteChg is the value change Event for Byte.
type ByteChg struct {
	eventBase
	ov, nv byte
}

func (ce ByteChg) String() string {
	return fmt.Sprintf("Chg byte [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce ByteChg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce ByteChg) NewValue() interface{} { return ce.nv }

// Rune is an Observable rune value. Rune implements change.Rune.
type Rune struct {
	v chgv.Rune
	ObservableBase
}

// NewRune creates a change-detectable Rune and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewRune(init rune, defaultTag interface{}, defaultChg change.Flags) *Rune {
	res := &Rune{v: chgv.Rune(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current rune value.
func (cv Rune) Get() rune { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// RuneChg is the value change Event for Rune.
type RuneChg struct {
	eventBase
	ov, nv rune
}

func (ce RuneChg) String() string {
	return fmt.Sprintf("Chg rune [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce RuneChg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce RuneChg) NewValue() interface{} { return ce.nv }

// Uint is an Observable uint value. Uint implements change.Uint.
type Uint struct {
	v chgv.Uint
	ObservableBase
}

// NewUint creates a change-detectable Uint and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewUint(init uint, defaultTag interface{}, defaultChg change.Flags) *Uint {
	res := &Uint{v: chgv.Uint(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current uint value.
func (cv Uint) Get() uint { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// UintChg is the value change Event for Uint.
type UintChg struct {
	eventBase
	ov, nv uint
}

func (ce UintChg) String() string {
	return fmt.Sprintf("Chg uint [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce UintChg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce UintChg) NewValue() interface{} { return ce.nv }

// Int is an Observable int value. Int implements change.Int.
type Int struct {
	v chgv.Int
	ObservableBase
}

// NewInt creates a change-detectable Int and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewInt(init int, defaultTag interface{}, defaultChg change.Flags) *Int {
	res := &Int{v: chgv.Int(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current int value.
func (cv Int) Get() int { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// IntChg is the value change Event for Int.
type IntChg struct {
	eventBase
	ov, nv int
}

func (ce IntChg) String() string {
	return fmt.Sprintf("Chg int [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce IntChg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce IntChg) NewValue() interface{} { return ce.nv }

// UintPtr is an Observable uintptr value. UintPtr implements change.UintPtr.
type UintPtr struct {
	v chgv.UintPtr
	ObservableBase
}

// NewUintPtr creates a change-detectable UintPtr and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewUintPtr(init uintptr, defaultTag interface{}, defaultChg change.Flags) *UintPtr {
	res := &UintPtr{v: chgv.UintPtr(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current uintptr value.
func (cv UintPtr) Get() uintptr { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// UintPtrChg is the value change Event for UintPtr.
type UintPtrChg struct {
	eventBase
	ov, nv uintptr
}

func (ce UintPtrChg) String() string {
	return fmt.Sprintf("Chg uintptr [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce UintPtrChg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce UintPtrChg) NewValue() interface{} { return ce.nv }

// String is an Observable string value. String implements change.String.
type String struct {
	v chgv.String
	ObservableBase
}

// NewString creates a change-detectable String and initializes its default
// tag and default change Flags.
// See also section Set Methods.
func NewString(init string, defaultTag interface{}, defaultChg change.Flags) *String {
	res := &String{v: chgv.String(init)}
	res.SetObsDefaults(defaultTag, defaultChg)
	return res
}

// Get returns the current string value.
func (cv String) Get() string { return cv.v.Get() }

// Set the value of cv to v. See also section Set Methods.
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
	cv.stub.update(e)
	return chg
}

// StringChg is the value change Event for String.
type StringChg struct {
	eventBase
	ov, nv string
}

func (ce StringChg) String() string {
	return fmt.Sprintf("Chg string [%v] -> [%v] flags=%x",
		ce.ov, ce.nv,
		ce.chg)
}

// Implement ValueChange
func (ce StringChg) OldValue() interface{} { return ce.ov }

// Implement ValueChange
func (ce StringChg) NewValue() interface{} { return ce.nv }
