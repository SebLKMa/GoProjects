package main

import ("fmt"; "os"; "path/filepath")

func listCurrentDir() {
    dir, err := os.Open(".")
    if err != nil {
        // TODO handle error
        return
    }
    defer dir.Close() // will be invoked when function exits
    
    fileInfo, err := dir.Readdir(-1) // -1 means all entries
    if err != nil {
        // TODO handle error
        return
    }
    
    for _, fi := range fileInfo {
        fmt.Println(fi.Name())
    } 
}

func listWalkCurrentDir() {
    filepath.Walk(".", 
        func(path string, info os.FileInfo, err error) error {
           fmt.Println(path)
           return nil
        },
    )
}

func main() {
    listCurrentDir()
    listWalkCurrentDir()
}