package yaml

import (
  "fmt"
  "time"
  "app/config"
)

type YamlFile struct {
  fileName string
  lastUpdated time.Time
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
