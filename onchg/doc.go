// Package onchg implements values with a simple function hook that
// handles value changes. First the hook is called before a change is
// made with pre=true. This decides if the change is blocked by,
// i.e. the hook returns 0. If the change is not blocked the hook
// provides the change flags for the Set method. The flags from the
// hook are overwritten if non-zero flags are passed to the value's
// Set method. If the hook is nil all values of this package behave
// exactly like their chgv counterparts.
package onchg
