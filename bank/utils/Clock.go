package utils

import "time"

type Clock interface {
	Now() time.Time
}
