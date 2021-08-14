package chgv

type Chg uint64

// Any tests if any of bits is set in c.
func (c Chg) Any(bits Chg) bool { return c&bits != 0 }

// All test if all bits are set in c.
func (c Chg) All(bits Chg) bool { return c&bits == bits }

// Map remaps bits of c to a new combination of bits.
//
// If the number of bits elements is odd then the last element is added to res
// if c is not zero. The even number of leading elements is treated as a list of
// pairs. The first element of a pair is checked to have any bits common with c.
// If so, the second element of the pair is added to res.
func (c Chg) Map(bits ...Chg) (res Chg) {
	l := len(bits) & ^1
	for i := 0; i < l; i++ {
		b := bits[i]
		i++
		if c.Any(b) {
			res |= bits[i]
		}
	}
	if l != len(bits) && c != 0 {
		res |= bits[l]
	}
	return res
}
