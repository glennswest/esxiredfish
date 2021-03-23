package main

// Reference
// https://www.supermicro.com/manuals/other/RedfishRefGuide.pdf

import "github.com/gin-gonic/gin"
import "fmt"
import "strconv"
ssh "github.com/glennswest/esxiredfish/sshclient"

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
                do_reset(resetCmd.ResetType,chassis);
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

func IsNumeric(s string) bool {
   _, err := strconv.ParseFloat(s, 64)
   return err == nil
}
 
func getvmid(thename string) string {
// vim-cmd vmsvc/getallvms
  

}

func do_reset(cmd string,chassis string){
/*
// "On",
//"ForceOff",
//"GracefulShutdown",
//"GracefulRestart",
//"ForceRestart",
//"Nmi",
//"ForceOn"
*/
       if (isNumeric(chassis) == false){
          chassis = getvmid(chassis);
          }

       switch cmd {
         case "On:
         case "ForceOff":
               
         case "GracefulSHutdown":
         case "GracefullRestart":
         case "Nmi":
         case "ForceOn":
         }

}

