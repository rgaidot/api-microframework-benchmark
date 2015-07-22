package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

type Route struct {
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)

  for _, route := range routes {
    router.
      Methods(route.Method).
      Path(route.Pattern).
      Handler(route.HandlerFunc)
  }

  return router
}

var routes = Routes{
  Route{
    "GET",
    "/api/v1/users",
    UsersIndex,
  },

  Route{
    "GET",
    "/api/v1/users/{userId}",
    UserShow,
  },
}
