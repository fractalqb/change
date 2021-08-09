package chg8

type Chg uint8

func (cs Chg) Then(chg Chg) Chg { return cs | chg }
