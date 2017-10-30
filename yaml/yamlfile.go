package yaml

import (
  "bytes"
  "fmt"
  "time"
  "app/config"
  "os/exec"
  "github.com/julienschmidt/httprouter"
)

type YamlFile struct {
  fileName string
  lastUpdated time.Time
}

type Log struct {
  fileName string
  logDate time.Time
  logContent string
}

func AllYamls() ([]YamlFile, error) {
  rows, err := config.DB.Query("SELECT * FROM yamls;")
  if err != nil {
    fmt.Println(err)
    return nil, err
  }
  defer rows.Close()

  yFiles := make([]YamlFile, 0)
  for rows.Next() {
    yFile := YamlFile{}
    err := rows.Scan(&yFile.fileName, &yFile.lastUpdated) // order matters
    if err != nil {
      fmt.Println(err)
      return nil, err
    }
    yFiles = append(yFiles, yFile)
  }

  if err = rows.Err(); err != nil {
    fmt.Println(err)
    return nil, err
  }

  return yFiles, nil
}

func PutYaml(ps httprouter.Params) (Log, error) {
  log := Log{}
  var out bytes.Buffer
  var err error
  // cmd := exec.Command("ls")
  cmd := exec.Command("./artillery/bin/artillery run " + ps.ByName("id"))
  cmd.Stdout = &out
  err = cmd.Run()
  if err != nil {
    fmt.Println(err)
    return log, err
  }
  log.logContent = out.String()
  log.logDate = time.Now()

  _, err = config.DB.Query("INSERT INTO logs (fileName, logDate, logContent) VALUES ($1, $2, $3);", ps.ByName("id"), log.logContent, log.logDate)
  if err != nil {
    fmt.Println(err)
    return log, err
  }

  return log, nil
}
