package yaml

import (
  "fmt"
  "time"
  "app/config"
  "errors"
  "net/http"
)

type YamlFile struct {
  fileName string
  fileNameOnDisk string
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
    err := rows.Scan(&yFile.fileName, &yFile.fileNameOnDisk, &yFile.lastUpdated) // order matters
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

func PutYaml(req *http.Request) (YamlFile, error) {
  yFile := YamlFile{}

  f, h, err := req.FormFile("q")
	if err != nil {
		return yFile, err
	}
	defer f.Close()

	// for your information
	fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)
  yFile.fileName = h.Filename
  yFile.lastUpdated = time.Now()
  yFile.fileNameOnDisk = "tmp"

  // insert values
	_, err = config.DB.Exec("INSERT INTO yamls (fileName, fileNameOnDisk, lastUpdated) VALUES ($1, $2, $3)", yFile.fileName, yFile.fileNameOnDisk, yFile.lastUpdated)
	if err != nil {
		return yFile, errors.New("500. Internal Server Error." + err.Error())
	}

	return yFile, nil
}
