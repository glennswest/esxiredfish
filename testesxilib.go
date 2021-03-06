package main
  
import (
        "fmt"
        esxi "github.com/glennswest/esxiredfish/esxilib"
)


func main() {
     dnsvmid := esxi.GetVmid("dns.gw.lo");
     fmt.Printf("Value for dns.gw.lo = %s\n",dnsvmid);
     devvmid := esxi.GetVmid("dev.gw.lo");
     fmt.Printf("Value for dev.gw.lo = %s\n",devvmid);

     dnspowerstate := esxi.GetPowerState("dns.gw.lo");
     fmt.Printf("Power for dns.gw.lo = %s\n",dnspowerstate);

     thelist := esxi.GetVmList();
     fmt.Printf("%v",thelist);

     esxi.PowerOffVm("master-0.bm.lo");
     result := esxi.GetPowerState("master-0.bm.lo");
     fmt.Printf("Powerstate %s (Expected off)\n",result);

}
     


