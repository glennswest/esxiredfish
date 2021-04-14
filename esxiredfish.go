package main

// Reference
// https://www.supermicro.com/manuals/other/RedfishRefGuide.pdf

import "os"
import "io"
import "github.com/gin-gonic/gin"
import "fmt"
import "strconv"
//import "time"
import "strings"
//import "net/http"
import "github.com/tidwall/gjson"
import "github.com/tidwall/sjson"
import "github.com/kardianos/service"
import esxi "github.com/glennswest/esxiredfish/esxilib"
//import "github.com/jinzhu/configor"
import "log"

type ResetCommand struct {
    ResetType string `json:"resettype"`
}



var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
        redfishserver();
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func init() {
}

func main() {

	svcConfig := &service.Config{
		Name:        "redfishesxi",
		DisplayName: "RedFishESXI",
		Description: "RedFish to VmWare ESXI BMC Server",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

func redfishserver() {
        //gin.SetMode(gin.ReleaseMode)
        gin.DisableConsoleColor();
        f, _ := os.Create("/var/log/esxiredfish.log");
        gin.DefaultWriter = io.MultiWriter(f);

        r := gin.New();
        r.Use(gin.Recovery());

        r.GET("/redfish/v1/Systems/:chassis", func(c *gin.Context){
                chassis := c.Param("chassis")
                fmt.Printf("Chassis: %v\n",chassis);
                sysinfo := GetSystemInfoBase(chassis)
                c.Data(200, "application/json", []byte(sysinfo))
                })
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
        r.GET("/redfish/v1/Actions", func(c *gin.Context){
              actions := "{}";
              c.Data(200, "application/json", []byte(actions))
              })
        r.GET("/redfish/v1/Systems/", func(c *gin.Context){
              systems := SetBaseSystemsJson();
	      c.Data(200, "application/json", []byte(systems))
              })
        r.GET("/redfish/v1/Systems", func(c *gin.Context){
              systems := SetBaseSystemsJson();
	      c.Data(200, "application/json", []byte(systems))
              })
        r.GET("/redfish/v1/", func(c *gin.Context){
              baseapi := GetBaseJson();
              c.Data(200, "application/json", []byte(baseapi))
              })
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func IsNumeric(s string) bool {
   _, err := strconv.ParseFloat(s, 64)
   return err == nil
}
 
func GetBaseJson() string {
json :=
  `{
    "@odata.context": "/redfish/v1/$metadata#ServiceRoot",
    "@odata.type": "#ServiceRoot.v1_1_1.ServiceRoot",
    "@odata.id": "/redfish/v1",
    "Id": "v1",
    "Name": "Root Service",
    "RedfishVersion": "1.0.1",
    "UUID": "ec5a4deb-9cdb-461e-8c25-5fdb22dd36d5",
    "Chassis": {
        "@odata.id": "/redfish/v1/Chassis"
    },
    "Systems": {
        "@odata.id": "/redfish/v1/Systems"
    },
    "Managers": {
        "@odata.id": "/redfish/v1/Managers"
    },
    "Tasks": {
        "@odata.id": "/redfish/v1/TaskService"
    },
    "SessionService": {
        "@odata.id": "/redfish/v1/SessionService"
    },
    "AccountService": {
        "@odata.id": "/redfish/v1/AccountService"
    },
    "EventService": {
        "@odata.id": "/redfish/v1/EventService"
    },
    "Registries": {
        "@odata.id": "/redfish/v1/Registries"
    },
    "CompositionService": {
        "@odata.id": "/redfish/v1/CompositionService"
    }
   }`;
  return json;
}
func SetBaseSystemsJson() string {
json := 
  `{
    "@odata.type": "#ComputerSystemCollection.ComputerSystemCollection",
    "Name": "Computer System Collection",
    "Id" : "Systems",
    "Members@odata.count": 0,
    "Members": [
    ],
    "@odata.context": "/redfish/v1/$metadata#Systems",
    "@odata.id": "/redfish/v1/Systems"
    }`

    thelist := esxi.GetVmList();
    println(thelist);
    value :=  gjson.Get(thelist,"vmcount").String()
    thecount, _ := strconv.Atoi(value);
    for i := 0; i < thecount; i++ {
       evalue := gjson.Get(thelist,"vmlist." + strconv.Itoa(i)).String();
       thename := gjson.Get(evalue,"name").String();
       odata := `{"@odata.id": "/redfish/v1/Systems/` + thename + `"}`;
       json, _ = sjson.SetRaw(json,"Members.-1",odata);
       }
   json, _ = sjson.Set(json,"Members@odata.count", thecount);
   return json
}

func GetSystemInfoBase(thesysid string) string {
json :=
 `{
    "@odata.type": "#ComputerSystem.v1_1_0.ComputerSystem",
    "Id": "${sysid}",
    "Name": "",
    "SystemType": "Physical", 
    "AssetTag": "${sysid}",
    "Manufacturer": "NCC Inc",
    "Model": "3500RX",
    "SKU": "8675309",
    "SerialNumber": "",
    "PartNumber": "",
    "Description": "",
    "UUID": "",
    "HostName": "${sysid}",
    "Status": {
        "State": "Enabled",
        "Health": "OK",
        "HealthRollUp": "OK"
    },
    "IndicatorLED": "Off",
    "PowerState": "On",
    "Boot": {
        "BootSourceOverrideEnabled": "Once",
        "BootSourceOverrideTarget": "Pxe",
        "BootSourceOverrideTarget@Redfish.AllowableValues": [
            "None",
            "Pxe",
            "Cd",
            "Usb",
            "Hdd",
            "BiosSetup",
            "Utilities",
            "Diags",
            "SDCard",
            "UefiTarget"
        ],
        "BootSourceOverrideMode": "UEFI",
        "UefiTargetBootSourceOverride": "/0x31/0x33/0x01/0x01"
    },
    "TrustedModules": [
        {
            "FirmwareVersion": "1.13b",
            "InterfaceType": "TPM1_2",
            "Status": {
                "State": "Enabled",
                "Health": "OK"
            }
        }
    ],
    "Oem": {
        "Contoso": {
            "@odata.type": "http://Contoso.com/Schema#Contoso.ComputerSystem",
            "ProductionLocation": {
                "FacilityName": "PacWest Production Facility",
                "Country": "USA"
            }
        },
        "Chipwise": {
            "@odata.type": "http://Chipwise.com/Schema#Chipwise.ComputerSystem",
            "Style": "Executive"
        }
    },
    "BiosVersion": "P79 v1.33 (02/28/2015)",
    "ProcessorSummary": {
        "Count": 2,
        "ProcessorFamily": "Multi-Core Intel(R) Xeon(R) processor 7xxx Series",
        "Status": {
            "State": "Enabled",
            "Health": "OK",
            "HealthRollUp": "OK"
        }
    },
    "MemorySummary": {
        "TotalSystemMemoryGiB": 96,
        "Status": {
            "State": "Enabled",
            "Health": "OK",
            "HealthRollUp": "OK"
        }
    },
    "Bios": {
        "@odata.id": "/redfish/v1/Systems/${sysid}/BIOS"
    },
    "Processors": {
        "@odata.id": "/redfish/v1/Systems/${sysid}/Processors"
    },
    "Memory": {
        "@odata.id": "/redfish/v1/Systems/${sysid}/Memory"
    },
    "EthernetInterfaces": {
        "@odata.id": "/redfish/v1/Systems/${sysid}/EthernetInterfaces"
    },
    "SimpleStorage": {
        "@odata.id": "/redfish/v1/Systems/${sysid}/SimpleStorage"
    },
    "LogServices": {
        "@odata.id": "/redfish/v1/Systems/${sysid}/LogServices"
    },
    "Links": {
        "Chassis": [
            {
                "@odata.id": "/redfish/v1/Chassis/1U"
            }
        ],
        "ManagedBy": [
            {
                "@odata.id": "/redfish/v1/Managers/BMC"
            }
        ]
    },
    "Actions": {
        "#ComputerSystem.Reset": {
            "target": "/redfish/v1/Systems/${sysid}/Actions/ComputerSystem.Reset",
            "ResetType@Redfish.AllowableValues": [
                "On",
                "ForceOff",
                "GracefulShutdown",
                "GracefulRestart",
                "ForceRestart",
                "Nmi",
                "ForceOn",
                "PushPowerButton"
            ]
        },
        "Oem": {
            "#Contoso.Reset": {
                "target": "/redfish/v1/Systems/${sysid}/Oem/Contoso/Actions/Contoso.Reset"
            }
        }
    },
    "@odata.context": "/redfish/v1/$metadata#ComputerSystem.ComputerSystem",
    "@odata.id": "/redfish/v1/Systems/${sysid}"
    }`



    powerstate := esxi.GetPowerState(thesysid);
    
    json, _ = sjson.SetRaw(json,"PowerState",powerstate);
    json = strings.Replace(json,"${sysid}",thesysid,-1);
    return json;
}

func getvmid(thename string) string {
/*
// vim-cmd vmsvc/getallvms
*/
     return "";

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
       switch cmd {
         case "On":
              esxi.PowerOnVm(chassis);
              break;
         case "ForceOff":
              esxi.PowerOffVm(chassis);
              break;
         case "GracefulShutdown":
              esxi.PowerShutdownVm(chassis);
              break;
         case "GracefullRestart":
              esxi.RestartVm(chassis);
              break;
         case "Nmi":
              esxi.PowerOffVm(chassis);
              break;
         case "ForceOn":
              esxi.PowerOnVm(chassis);
              break;
         }

}

