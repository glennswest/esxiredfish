package main
  
import (
        "fmt"
        "log"
        "io/ioutil"
        "golang.org/x/crypto/ssh"
        "os"
)


func main() {
     user := "root";
     host := "192.168.1.150";
     cmd := "ls";
     output, error := SshClientCmd(user,host,cmd);
     fmt.Println("result = %s\n",output);
}
     


