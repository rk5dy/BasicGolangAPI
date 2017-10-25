package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
  "app/yaml"
)

func main() {
  router := httprouter.New()
  router.GET("/", yaml.Index)
  http.ListenAndServe(":8080", router);
}
