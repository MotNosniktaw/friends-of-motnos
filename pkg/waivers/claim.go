package waivers

import "gorm.io/gorm"

type Claim struct {
	gorm.Model
	TeamId string
	Amount int32
}
