package main

import ("fmt"; "encoding/gob"; "net")

func server() {
    // listen on a port
    listener, err:= net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    for {
        // accept a connection
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            continue    
        }
        // handle connection
        go handleServerConnection(conn)
    }
}

func handleServerConnection(conn net.Conn) {
    // receive the message
    var msg string
    err := gob.NewDecoder(conn).Decode(&msg) // msg decoder
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Received<<", msg)
    }
    conn.Close()
}

func client() {
    // connect to server
    conn, connErr := net.Dial("tcp", "127.0.0.1:9999")
    if connErr != nil {
        fmt.Println(connErr)
        return
    }   

    // send the messsage
    msg := "hello go tcp"
    fmt.Println("Sending>>>", msg)
    sendError := gob.NewEncoder(conn).Encode(msg) // msg encoder
    if sendError != nil {
        fmt.Println(sendError)
    }
    
    conn.Close()
}

func main() {
    go server() // start server
    go client() // start client
    
    var input string
    fmt.Scanln(&input) // wait on main thread
}