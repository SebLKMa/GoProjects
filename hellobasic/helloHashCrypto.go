package main

import ("fmt"; "hash/crc32"; "os"; "io"; "crypto/sha1")

func f1checksum() {
    // create a hasher, crc32 computes a 32-bit hash
    h := crc32.NewIEEE()
    
    // write binary data to hasher
    s := "hallo"
    h.Write([]byte(s))
    
    // calculate the crc32 checksum
    v1 := h.Sum32()

    fmt.Println(s, v1)

    // create another hasher to hash the same data
    h2 := crc32.NewIEEE()
    h2.Write([]byte(s))
    v2 := h.Sum32()
    
    // checksum v1 and v2 should be the same
    fmt.Println(s, v2)    
}

func getHash(filename string) (uint32, error) {
    // open the file
    file, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    defer file.Close()
    
    // create a hasher
    hasher := crc32.NewIEEE()
    
    // copy the file to the hasher
    // - copy takes (dst, src) and returns (bytesWritten, error)
    _, err = io.Copy(hasher, file)
    // ignore bytes writtenm just handle the error
    if err != nil {
        return 0, err
    }
    
    return hasher.Sum32(), nil
}

func f2checksumfile() {
    // test.txt is created from helloFileIO.go
    
    h1, err := getHash("test.txt")
    if err != nil {
        return
    }
    h2, err := getHash("test.txt")
    if err != nil {
        return
    }
    fmt.Println(h1, h2, h1 == h2)
}

func f3onewayHash() {
    hasher := sha1.New() // sha1 computes a 160-bit hash
    hasher.Write([]byte("hallo"))
    bytes := hasher.Sum([]byte{}) // there is no native type for 160-bit number, a slice of 20 bytes is used as return value
    fmt.Println(bytes)
}

func main() {
    f1checksum()
    f2checksumfile()
    f3onewayHash()
}