package models

import "time"

type Travel struct {
	ID          int64
	Destination string
	Date        time.Time
	AuxDate     string
	Budget      float64
	Clothes     Clothes
}

type Clothes struct {
	ID        int64
	Underwear uint8
	Pants     uint8
	Shirts    uint8
	TShirts   uint8
	Shoes     uint8
}
