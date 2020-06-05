package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
    s := fmt.Sprintf("%v", ip[0])

    for _, v := range ip[1:] {
        s += fmt.Sprintf(".%v", v)
    }
    return s
}

func main() {
    
    hosts := map[string]IPAddr {
        "loopback": {127,0,0,1},
        "googleDNS": {8,8,8,8},
    }

    for name, ip := range hosts {
        fmt.Printf("%v : %v\n", name, ip)
    }
}
