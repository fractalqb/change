// Package change has the common parts for the sub packages that help
// to track changes of values in an application. The sub packages
// implement different trade offs between resource consumption –
// memory, time – and features richness.  The changeable value types
// of the sub packages all implement their respective interface in
// this package which makes migration between the sub packages more
// easy.
package change
