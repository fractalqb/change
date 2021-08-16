// Package chgv implements values that track changes with a simple
// flag set of type 'change.Flags'. The flag(s) passe to any Set call
// are retuned if the underlying values changed. Otherwiese 0 is
// returned. Note that if one passes 0 as flag to Set the returned
// value is rather useless. However this will not affect the actual
// value change.
//
// While these changeable values are rather bare bones they come
// without memory overhead – unlike most observable libraries.
package chgv
