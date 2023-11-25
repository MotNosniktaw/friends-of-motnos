package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/motnosniktaw/nflfantasy/internal/db"
	"github.com/motnosniktaw/nflfantasy/pkg/models"
)

func main() {
	r := gin.Default()

	db.Connect()

	players := r.Group("/players")
	players.GET("/", func(ctx *gin.Context) {
		var players []models.Player
		db.DB.Preload("nf;_teams").Find(&players)

		ctx.JSON(http.StatusOK, players)
	})

	http.ListenAndServe(":50052", r)
}
