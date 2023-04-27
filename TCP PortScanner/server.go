package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()
    fmt.Println("Server started. Listening on port 8080.")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    fmt.Println("New client connected:", conn.RemoteAddr().String())

    openPorts := scanPorts()
    _, err := conn.Write([]byte(openPorts))
    if err != nil {
        fmt.Println("Error sending open ports to client:", err)
        return
    }
    fmt.Println("Open ports sent to client.")
}

func scanPorts() string {
    host := "example.com"
    startPort := 1
    endPort := 200

    var openPorts string

    for port := startPort; port <= endPort; port++ {
        address := fmt.Sprintf("%s:%d", host, port)
        conn, err := net.DialTimeout("tcp", address, 100*time.Millisecond)
        if err == nil {
            conn.Close()
            openPorts += fmt.Sprintf("%d\n", port)
        }
    }

    return openPorts
}

