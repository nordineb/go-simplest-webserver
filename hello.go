package main

import (
  "net/http"
  "fmt"
  "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello world\n"))
    })

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
      fmt.Printf("serving: %v\n", err)
      os.Exit(1)
    }
}
