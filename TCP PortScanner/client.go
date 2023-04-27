package main

import (
    "bufio"
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()
    fmt.Println("Connected to server.")

    response, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println("Error reading server response:", err)
        return
    }
    fmt.Println("Open ports:")
    fmt.Print(response)
}

