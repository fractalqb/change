package change

import (
	"time"
)

type Time interface {
	Get() time.Time
	Set(v time.Time, chg Flags) Flags
}
