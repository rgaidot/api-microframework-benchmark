package main

import (
  "net/http"
  "encoding/json"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  users := Users{
    User{Id: "a1de359e-6bf3-4394-bf88-9de1397f50c1", FirstName: "Régis", LastName: "Gaidot"},
    User{Id: "843693e6-6e63-43ed-8dad-e504410d9d6c", FirstName: "Ethan", LastName: "Gaidot"},
  }

  json.NewEncoder(w).Encode(users)
}

func UserShow(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  user := User{Id: "a1de359e-6bf3-4394-bf88-9de1397f50c1", FirstName: "Régis", LastName: "Gaidot"}
  json.NewEncoder(w).Encode(user)
}
