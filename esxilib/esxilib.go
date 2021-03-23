package esxilib

import (
        "strings"
        "github.com/jinzhu/configor"
        ssh "github.com/glennswest/esxiredfish/sshclient"
)


var Config = struct {
    Host string `default:"192.168.1.150"`
    User string `default:"root"`
}{}

func init(){
     configor.Load(&Config,"config.yaml");
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

func GetPowerState(thevmid string) string {
// vim-cmd vmsvc/power.getstate <Vmid>

     cmd := "vim-cmd vmsvc/power.getstate " + thevmid;
     result, _ := doCmd(cmd);
     lines := strings.Split(result,"\n");
     if (strings.Contains(lines[0],"vim.fault.NotFound")){
        return("invalid");
        }
     switch(lines[1]){
         case "Powered on": return("on");
         case "Powered off": return("off");
         default:      return("unknown");
         }
     return "impossible";
}

func doCmd(cmd string) (string, error){
     out, err := ssh.SshClientCmd(Config.User,Config.Host, cmd);
     return out, err;
}

