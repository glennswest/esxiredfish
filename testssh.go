package main
  
import (
        "fmt"
        ssh "github.com/glennswest/esxiredfish/sshclient"
)


func main() {
     user := "root";
     host := "192.168.1.150:22";
     cmd := "ls";
     output, _ := ssh.SshClientCmd(user,host,cmd);
     fmt.Println("result = %s\n",output);
}
     


