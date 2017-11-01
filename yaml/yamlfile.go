package yaml

import (
  "bytes"
  "fmt"
  "time"
  "io"
  "app/config"
  "errors"
  "mime/multipart"
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
  yFile.lastUpdated = time.Now()

  //get file
  f, h, err := req.FormFile("yaml")
	if err != nil {
		return yFile, err
	}
  yFile.fileName = h.Filename

	defer f.Close()
  var part io.Writer
  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  part, err = writer.CreateFormFile("file", yFile.fileName)
  if err != nil {
      return yFile, err
  }
  _, err = io.Copy(part, f)
  err = writer.Close()

  request, err := http.NewRequest("POST", "localhost:5000/upload", body)
  request.Header.Set("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, err := client.Do(request)
  if err != nil {
      return yFile, err
  } else {
    body := &bytes.Buffer{}
    _, err := body.ReadFrom(resp.Body)
    if err != nil {
      return yFile, err
    }
    resp.Body.Close()
    yFile.fileNameOnDisk = body.String()
  }
  // insert values
	_, err = config.DB.Exec("INSERT INTO yamls (fileName, fileNameOnDisk, lastUpdated) VALUES ($1, $2, $3)", yFile.fileName, yFile.fileNameOnDisk, yFile.lastUpdated)
	if err != nil {
		return yFile, errors.New("500. Internal Server Error." + err.Error())
	}

	return yFile, nil
}
