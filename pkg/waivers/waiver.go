package waivers

import "gorm.io/gorm"

type Waiver struct {
	gorm.Model
	PlayerId string
	ClosesAt int64
}
