package yaml

import (
  "fmt"
  "net/http"
  "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  yamls, err := AllYamls()
  if err != nil {
    http.Error(w, http.StatusText(500), http.StatusInternalServerError)
    return
  }
  for _,yaml := range yamls {
    fmt.Fprintf(w, "%s %s", yaml.fileName, yaml.lastUpdated.String())
  }
}
