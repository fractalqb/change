define(`onchgtype', `// $2 tracks changes of $1 values with caller-provided Flags.
// See also chgv package documentation.
type $2 $1

func (cv $2) Get() $1 { return $1(cv) }

func (cv *$2) Set(v $1, chg change.Flags) change.Flags {
	if $1(*cv) == v {
		return 0
	}
	*cv = $2(v)
	return chg
}
')dnl
dnl
package chgv

import "git.fractalqb.de/fractalqb/change"

onchgtype(`bool', `Bool')
onchgtype(`uint8', `Uint8')
onchgtype(`uint16', `Uint16')
onchgtype(`uint32', `Uint32')
onchgtype(`uint64', `Uint64')
onchgtype(`int8', `Int8')
onchgtype(`int16', `Int16')
onchgtype(`int32', `Int32')
onchgtype(`int64', `Int64')
onchgtype(`float32', `Float32')
onchgtype(`float64', `Float64')
onchgtype(`complex64', `Complex64')
onchgtype(`complex128', `Complex128')
onchgtype(`byte', `Byte')
onchgtype(`rune', `Rune')
onchgtype(`uint', `Uint')
onchgtype(`int', `Int')
onchgtype(`uintptr', `UintPtr')
onchgtype(`string', `String')
