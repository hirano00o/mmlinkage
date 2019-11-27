package adapters

import (
  "fmt"
  "net/http"
  "bytes"
  "errors"
)

type MattermostAdapter struct {}

func (a *MattermostAdapter) Post(url string, message []byte) (err error) {
  resp, err := http.Post(url, "application/json", bytes.NewBuffer(message))
  if err != nil {
    return
  }

  if resp.StatusCode != http.StatusOK {
    errors.New(fmt.Sprintf("StatusCode is %d", resp.StatusCode))
  }

  return
}
