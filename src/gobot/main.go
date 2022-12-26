package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
)

// main function to serve bot app
func main() {
    port      := os.Getenv("PORT")
    fmt.Printf("listening on port \n"+port)
    http.HandleFunc("/", BotHandler)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
