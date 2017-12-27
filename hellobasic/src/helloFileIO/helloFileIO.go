package main

import ("fmt"; "os"; "io/ioutil")

const filename string = "test.txt"

func f1create() {
    newFile, err := os.Create(filename)
    if err != nil {
        // TODO handle error
        return
    }
    defer newFile.Close() // will be invoked when function exits
    
    newFile.WriteString("To err is human")
}

func f2read() {
    myFile, err := os.Open(filename)
    if err != nil {
        // TODO handle error
        return
    }
    defer myFile.Close()
    
    // get file size
    myStat, err := myFile.Stat()
    if err != nil {
        // TODO handle error
        return
    }
    
    // read the file
    bytes := make([]byte, myStat.Size())
    _, err = myFile.Read(bytes)
    if err != nil {
        // TODO handle error
        return
    }
    
    str := string(bytes)
    fmt.Println(str)
}

func f3read() {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        // TODO handle error
        return
    }
    
    str := string(bytes)
    fmt.Println(str)
}

func main() {
    f1create()
    f2read()
    f3read()
}