package daten

import "time"

type Bankdata struct {
	Iban     string
	Amount   float32
	Balance  float32
	BooketAt time.Time
}