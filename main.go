package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/atopos31/tiny-url/db"
	"github.com/atopos31/tiny-url/models"
	"github.com/atopos31/tiny-url/utils"
	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

var domain string

func init() {
	domain = os.Getenv("SERVER_DOMAIN")
}

func main() {
	router := gin.Default()

	router.POST("/create", handlerCreate)
	router.GET("/:url", handlerRedirect)

	port := os.Getenv("SERVER_PORT")
	router.Run(":" + port)
}

func handlerCreate(c *gin.Context) {
	req := new(models.ReqCreateTinyURL)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, models.Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	tinyURL := utils.GenerateTinyURL(req.URL)

	if err := db.Rdb.Set(c, tinyURL, req.URL, time.Duration(req.Time)*time.Hour*24).Err(); err != nil {
		c.JSON(http.StatusOK, models.Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Code: 200,
		Data: fmt.Sprintf("%s/%s", domain, tinyURL),
	})
}

func handlerRedirect(c *gin.Context) {
	tinyURL := c.Param("url")
	if strings.EqualFold(tinyURL, "") {
		c.JSON(http.StatusOK, models.Response{
			Code:    400,
			Message: "Invalid tinyURL",
		})
		return
	}
	URL, err := db.Rdb.Get(c, tinyURL).Result()
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			Code:    404,
			Message: "Redirect error: " + err.Error(),
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, URL)
}
