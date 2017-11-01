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
    fmt.Fprintf(w, "%s %s %s", yaml.fileName, yaml.fileNameOnDisk, yaml.lastUpdated.String())
  }
}

func Upload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  yaml, err := PutYaml(r)
  if err != nil {
    http.Error(w, http.StatusText(500), http.StatusInternalServerError)
    return
  }
  fmt.Println(yaml.fileName)
}
