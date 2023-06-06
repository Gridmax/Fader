package main


import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/Gridmax/Fader/utility/configload"
)

func main() {
  r := gin.Default()

  config, _ := configload.LoadConfig("config.yaml")

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  r.Run(":"+config.ServicePort)
}

