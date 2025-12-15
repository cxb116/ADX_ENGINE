package regService

import (
	"encoding/json"
	"github.com/cxb116/ADX_ENGINE/registerEngine/registry"

	"log"
	"net/http"

	"sync"
)

const ServerPort = ":80"
const ServicesUrl = "http://127.0.0.1:80"

type register struct {
	registrations []registry.Registration
	mutex         *sync.RWMutex
}

func (this *register) add(reg registry.Registration) error {
	this.mutex.Lock()
	this.registrations = append(this.registrations, reg)
	this.mutex.Unlock()
	return nil
}

var regs = register{
	registrations: make([]registry.Registration, 0),
	mutex:         new(sync.RWMutex),
}

type RegisterService struct{}

func (res RegisterService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var registration registry.Registration
		err := decode.Decode(&registration)
		if err != nil {
			log.Println("decode err:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf(
			"Adding Service: %s with URL: %s",
			registration.ServiceName,
			registration.ServiceURL,
		)

		if err := regs.add(registration); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
