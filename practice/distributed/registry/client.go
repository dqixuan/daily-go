package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewDecoder(buf)
	err := enc.Decode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServiceUrl, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service. Registry service responsed with code %v", res.StatusCode)
	}
	return nil
}
