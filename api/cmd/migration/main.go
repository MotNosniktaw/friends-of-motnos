package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/motnosniktaw/nflfantasy/internal/db"
	"github.com/motnosniktaw/nflfantasy/internal/stats"
	"github.com/motnosniktaw/nflfantasy/pkg/models"
)

func main() {
	godotenv.Load()

	db.Connect()

	fmt.Println("Connected to database", db.DB)

	db.DB.AutoMigrate(&models.Player{})
	db.DB.AutoMigrate(&models.NFLTeam{})

	for _, teamCode := range models.GetAllTeamCodes() {
		var team models.NFLTeam
		db.DB.Where("code = ?", teamCode).FirstOrCreate(&team, models.NFLTeam{Code: teamCode})
	}

	statsClientApiKey := os.Getenv("STATS_CLIENT_API_KEY")
	fmt.Println("Using stats client api key", statsClientApiKey)
	statsClient := stats.NewClient(statsClientApiKey)

	players, err := statsClient.GetPlayers()
	if err != nil {
		panic(err)
	}

	for _, player := range players {
		var playerModel models.Player
		db.DB.Where("name = ?", player.Name).FirstOrCreate(&playerModel, models.Player{
			Name:     player.Name,
			TeamCode: models.TeamCode(player.TeamCode),
			Position: models.Position(player.Position),
		})
	}
}
