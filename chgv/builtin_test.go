package chgv

import (
	"git.fractalqb.de/fractalqb/change"
)

var (
	_ change.Bool       = (*Bool)(nil)
	_ change.Uint8      = (*Uint8)(nil)
	_ change.Uint16     = (*Uint16)(nil)
	_ change.Uint32     = (*Uint32)(nil)
	_ change.Uint64     = (*Uint64)(nil)
	_ change.Int8       = (*Int8)(nil)
	_ change.Int16      = (*Int16)(nil)
	_ change.Int32      = (*Int32)(nil)
	_ change.Int64      = (*Int64)(nil)
	_ change.Float32    = (*Float32)(nil)
	_ change.Float64    = (*Float64)(nil)
	_ change.Complex64  = (*Complex64)(nil)
	_ change.Complex128 = (*Complex128)(nil)
	_ change.Byte       = (*Byte)(nil)
	_ change.Rune       = (*Rune)(nil)
	_ change.Uint       = (*Uint)(nil)
	_ change.Int        = (*Int)(nil)
	_ change.UintPtr    = (*UintPtr)(nil)
	_ change.String     = (*String)(nil)
)
