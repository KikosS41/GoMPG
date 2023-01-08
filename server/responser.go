package server

import (
	"encoding/json"
	"github.com/KikosS41/GoMPG/server/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func responser(players []entities.Player) {
	r := gin.Default()

	r.POST("/connect", func(c *gin.Context) {
		newPlayer := entities.Player{}
		err := c.BindJSON(&newPlayer)
		if err != nil {
			c.String(http.StatusBadRequest, "The player you want to generate is not compatible !")
			return
		}

		err = AddNewPlayer(&players, newPlayer)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		response, _ := json.Marshal(players)
		c.String(http.StatusAccepted, string(response))
	})

	r.POST("/update", func(c *gin.Context) {
		player := entities.Player{}
		err := c.BindJSON(&player)
		if err != nil {
			c.String(http.StatusBadRequest, "The player you want to generate is not compatible !")
			return
		}

		players, err = UpdatePlayer(players, player)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		response, _ := json.Marshal(players)
		c.String(http.StatusAccepted, string(response))
	})

	r.POST("/disconnect", func(c *gin.Context) {
		newPlayer := entities.Player{}
		err := c.BindJSON(&newPlayer)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		players, err = DeletePlayer(players, newPlayer)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		response, _ := json.Marshal(players)
		c.String(http.StatusAccepted, string(response))
	})

	r.GET("/infos", func(c *gin.Context) {
		response, _ := json.Marshal(players)
		c.String(http.StatusAccepted, string(response))
	})

	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}
