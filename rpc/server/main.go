package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func main() {
	r := gin.Default()
	r.GET("/tick", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{Data: time.Now().Format("2006-01-02 03:04:05")})
	})

	r.Run(":81")
}
