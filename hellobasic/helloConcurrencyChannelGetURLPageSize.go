package main

import ("fmt"; "io/ioutil"; "net/http")

type HomePageSize struct {
    URL string
    Size int
}

func main() {
    urls := []string{
        "http://www.apple.com",
        "http://www.amazon.com",
        "http://www.google.com",
        "http://www.microsoft.com",    
    }
    
    resultsChannel := make(chan HomePageSize)

    // for each url, start a goroutine to get url page size
    for _, url := range urls {
        go func(urlInput string) {
            res, err := http.Get(urlInput)
            if err != nil {
                panic(err)
            }
            defer res.Body.Close()
            
            bytes, err := ioutil.ReadAll(res.Body)
            if err != nil {
                panic(err)
            }
            resultsChannel <- HomePageSize{URL:urlInput, Size:len(bytes)} // send result to channel
            fmt.Println("Sent to channel", urlInput, len(bytes))
            }(url) // closure
    }
    
    var biggest HomePageSize
    
    // for each url, receive result from channel and store the biggest size url
    for range urls {
        result := <-resultsChannel
        if result.Size > biggest.Size {
            biggest = result
        }
    }
    
    fmt.Println("The biggest home page:", biggest.URL)
    fmt.Println("Size:", biggest.Size)
 }

