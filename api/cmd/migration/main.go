package main

import (
	"fmt"

	"github.com/motnosniktaw/nflfantasy/internal/db"
	"github.com/motnosniktaw/nflfantasy/pkg/models"
)

func main() {

	db.Connect()

	fmt.Println("Connected to database", db.DB)

	db.DB.AutoMigrate(&models.Player{})
	db.DB.AutoMigrate(&models.NFLTeam{})

	for _, teamCode := range models.GetAllTeamCodes() {
		var team models.NFLTeam
		db.DB.Where("code = ?", teamCode).FirstOrCreate(&team, models.NFLTeam{Code: teamCode})
	}

}
