package meterconv

import (
	"fmt"
)

type Meter float64
type Kilometer float64

func (m Meter) String() string {
	return fmt.Sprintf("%gM", m)
}

func (km Kilometer) String() string {
	return fmt.Sprintf("%gKM", km)
}
