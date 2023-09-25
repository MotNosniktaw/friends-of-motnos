package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/motnosniktaw/nflfantasy/internal/db"
	"github.com/motnosniktaw/nflfantasy/pkg/waivers"
	"gorm.io/gorm"
)

func main() {
	db := db.Connect()
	db.AutoMigrate(&waivers.Waiver{})
	db.AutoMigrate(&waivers.Claim{})

	r := gin.Default()
	r.GET("/player", func(c *gin.Context) {
		id := c.Query("id")
		fmt.Println(id)

		var waiver waivers.Waiver
		err := db.First(&waiver, "player_id = ?", id).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(204, "No record matching that id")
			return
		}

		c.JSON(200, waiver)
	})

	r.POST("/player", func(c *gin.Context) {
		decoder := json.NewDecoder(c.Request.Body)
		input := struct {
			Id string `json:"id"`
		}{}
		decoder.Decode(&input)

		fmt.Println(input)

		err := db.Create(&waivers.Waiver{PlayerId: input.Id,
			ClosesAt: time.Now().Add(time.Hour * 24).Unix()}).Error

		if err != nil {
			panic("Could not add record to database")
		}

		c.JSON(200, "Waiver added")
	})

	r.Run(":50052")
}
