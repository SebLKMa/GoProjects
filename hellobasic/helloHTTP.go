package main

import ("fmt"; "net/http"; "io")

func helloResponse(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content.Type", "text/html",)
    
    // use ` character for string with multiline
    // <head> appears as tab title
    // <body> appears in page body
    io.WriteString(
        response, 
        `<!DOCTYPE html>
        <html>
          <head>
            <title>Hello GO HTTP head</title>
          </head>
          <body>
            Hello GO HTTP Body
          </body>
        </html>`,
    )
}

func startHttpByFile() {
    http.Handle(
        "/assets/",
        http.StripPrefix(
        "/assets/",
        http.FileServer(http.Dir("assets")),
        ),
    )
}

func main() {
    http.HandleFunc("/hello", helloResponse) // handles URL route /hello by calling helloResponse
    
    fmt.Println("Starting HTTP server...")
    http.ListenAndServe(":9009", nil)
}