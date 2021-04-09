package esxilib

import (
        "fmt"
        "strings"
        "github.com/jinzhu/configor"
        ssh "github.com/glennswest/esxiredfish/sshclient"
)
//import "github.com/tidwall/gjson"
import "github.com/tidwall/sjson"


var Config = struct {
    Host string `default:"192.168.1.150"`
    User string `default:"root"`
}{}

func init(){
     configor.Load(&Config,"/etc/esxiredfish.yaml");
}

func GetVmList() string {
// vim-cmd vmsvc/getallvms
// 8      dns.gw.lo               [datastore1] dns.gw.lo/dns.gw.lo.vmx                           rhel8_64Guest           vmx-14  
     thelist := `{"vmcount": 0, "vmlist": []}`;
     result, _ := doCmd("vim-cmd vmsvc/getallvms"); 

     cnt := 0;
     lines := strings.Split(result,"\n");
     for _, s := range lines[1:] {
        values := strings.Fields(s);
        if (len(values) > 1){
           cnt = cnt + 1;
           element := `{}`;
           element, _ =  sjson.Set(element,"vmid",values[0]);
           element, _ =  sjson.Set(element,"name",values[1]);
           thelist, _ = sjson.SetRaw(thelist,"vmlist.-1",element); 
           }
        }
    thelist, _ = sjson.Set(thelist,"vmcount",cnt);
    return thelist;
}

func GetVmid(thename string) string {
// vim-cmd vmsvc/getallvms
// 8      dns.gw.lo               [datastore1] dns.gw.lo/dns.gw.lo.vmx                           rhel8_64Guest           vmx-14  
     result, _ := doCmd("vim-cmd vmsvc/getallvms"); 
     lines := strings.Split(result,"\n");
     for _, s := range lines {
         values := strings.Fields(s);
         if (strings.Contains(values[1],thename)){
            return(values[0]);
            }
         }
    return "";
}


func PowerOffVm(thevm string){
// vim-cmd vmsvc/power.off ${vmid}
     thevmid := GetVmid(thevm);
     cmd := "vim-cmd vmsvc/power.off " + thevmid;
     doCmd(cmd);
}

func PowerOnVm(thevm string){
// vim-cmd vmsvc/power.on ${vmid}
     thevmid := GetVmid(thevm);
     cmd := "vim-cmd vmsvc/power.on " + thevmid;
     fmt.Printf("PowerON: %s\n",thevmid);
     fmt.Printf("%s\n",cmd);
     doCmd(cmd);
}

func PowerShutdownVm(thevm string){
// vim-cmd vmsvc/power.shutdown ${vmid}

     thevmid := GetVmid(thevm);
     cmd := "vim-cmd vmsvc/power.shutdown " + thevmid;
     doCmd(cmd);
}

func RestartVm(thevm string){
// vim-cmd vmsvc/power.reboot ${vmid}

     thevmid := GetVmid(thevm);
     cmd := "vim-cmd vmsvc/power.reboot " + thevmid;
     doCmd(cmd);
}



func GetPowerState(thevm string) string {
// vim-cmd vmsvc/power.getstate <Vmid>

     thevmid := GetVmid(thevm);
     cmd := "vim-cmd vmsvc/power.getstate " + thevmid;
     result, _ := doCmd(cmd);
     lines := strings.Split(result,"\n");
     if (strings.Contains(lines[0],"vim.fault.NotFound")){
        return("invalid");
        }
     switch(lines[1]){
         case "Powered on": return("On");
         case "Powered off": return("Off");
         default:      return("unknown");
         }
     return "impossible";
}

func doCmd(cmd string) (string, error){
     out, err := ssh.SshClientCmd(Config.User,Config.Host, cmd);
     return out, err;
}

