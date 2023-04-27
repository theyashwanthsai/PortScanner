package main

import (
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

    message := "Hello, server!"
    _, err = conn.Write([]byte(message))
    if err != nil {
        fmt.Println("Error sending message to server:", err)
        return
    }
    fmt.Println("Message sent to server.")

    response := make([]byte, 4096)
    n, err := conn.Read(response)
    if err != nil {
        fmt.Println("Error reading server response:", err)
        return
    }
    fmt.Println("Response from server:", string(response[:n]))
}
