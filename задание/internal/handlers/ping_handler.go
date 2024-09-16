package handlers

import (
	"testovoeZadanir/internal/util"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	util.Logger.Info("Ping endpoint hit")
	c.String(200, "ok")
}