package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
  "app/yaml"
)

func main() {
  router := httprouter.New()
  router.GET("/", yaml.Index)
  router.GET("/yaml/:id", yaml.ArtilleryFire)

  http.ListenAndServe(":6060", router);
}
