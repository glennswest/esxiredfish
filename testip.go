package main

import (
    "fmt"
    "net"
)

func main() {
    addrs, _ := net.InterfaceAddrs()
    fmt.Printf("%v\n", addrs)
    for _, addr := range addrs {
        fmt.Println("IPv4: ", addr)
    }
}

