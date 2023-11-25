package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Position Position
	TeamCode TeamCode
	Team     NFLTeam `gorm:"foreignKey:TeamCode"`
}
