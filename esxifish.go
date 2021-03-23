package main

import "github.com/gin-gonic/gin"
import "fmt"

type ResetCommand struct {
    ResetType string `json:"resettype"`
}

func main() {
	r := gin.Default();
        r.POST("/redfish/v1/Systems/:chassis/Actions/ComputerSystem.Reset", func(c *gin.Context) {

                var resetCmd ResetCommand
                chassis := c.Param("chassis")
                c.BindJSON(&resetCmd);
                fmt.Printf("ResetType: %v\n",resetCmd.ResetType);
                fmt.Printf("Chassis: %v\n",chassis);
		c.JSON(200, gin.H{
			"message": "pong",
		})
        })
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

