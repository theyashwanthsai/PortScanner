package main

import (
    "fmt"
    "net"
    "time"
)

func scanPort(address string, port int, timeout time.Duration) bool {
    conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", address, port), timeout)
    if err != nil {
        return false
    }
    conn.Close()
    return true
}

func main() {
    address := "example.com"
    timeout := time.Second * 3
    for port := 1; port <= 65535; port++ {
        if scanPort(address, port, timeout) {
            fmt.Printf("Port %d is open\n", port)
        }
    }
}
