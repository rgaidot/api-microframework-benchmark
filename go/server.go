package main

import (
  "net/http"
  "log"
)

func main() {
  router := NewRouter()
  log.Print("http://0.0.0.0:3000/api/v1/users")
  log.Fatal(http.ListenAndServe(":3000", router))
}
