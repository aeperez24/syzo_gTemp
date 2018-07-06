package temperature

import (
	"time"
)

type Temperature struct {
	Date        time.Time
	Measurement float64
	Name        string
}
