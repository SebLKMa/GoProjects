package main

import ("fmt"; "net"; "net/rpc")

// RPC exposes methods that can be invoked by remote clients

// server methods
type Server struct {}
func (this *Server) Negate(i int64, reply *int64) error {
    *reply = -i // dereference to set reply's value
    return nil // no error
}

func server() {
    rpc.Register(new(Server))
    listener, err := net.Listen("tcp", ":9999")
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
        go rpc.ServeConn(conn)
    }
}

func client() {
    // connect to server
    conn, connErr := rpc.Dial("tcp", "127.0.0.1:9999")
    if connErr != nil {
        fmt.Println(connErr)
        return
    }   

    // send the messsage
    msg := int64(42)
    var result int64
    fmt.Println("Sending>>>", msg)
    sendError := conn.Call("Server.Negate", msg, &result)
    if sendError != nil {
        fmt.Println(sendError)
    } else {
        fmt.Println("Server.Negate =", result)
    }
    
    conn.Close()
}

func main() {
    go server() // start server
    go client() // start client
    
    var input string
    fmt.Scanln(&input) // wait on main thread
}