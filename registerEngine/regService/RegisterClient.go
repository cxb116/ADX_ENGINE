package regService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cxb116/ADX_ENGINE/registerEngine/registry"
	"net/http"
)

func RegisterClient(reg registry.Registration) error {

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(reg)
	if err != nil {
		return err
	}

	res, err := http.Post(ServicesUrl, "application/json", buf)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service . Registry service"+
			"responded with code: %v", res.StatusCode)
	}
	return nil
}
