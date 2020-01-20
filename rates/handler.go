package rates

import (
	"encoding/json"
	"net/http"
)

//HandlerInterface interface
type HandlerInterface interface {
	GetRates(w http.ResponseWriter, r *http.Request)
	SyncData()
}

//Handler struct
type Handler struct {
	Uc UsecaseInterface
}

//NewHandler function
func NewHandler(uc UsecaseInterface) HandlerInterface {
	return &Handler{uc}
}

//GetRates function
func (hd *Handler) GetRates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	prefix := r.URL.Path
	result, err := hd.Uc.GetRates(prefix)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(result)
}

//SyncData functions
func (hd *Handler) SyncData() {
	envelop := hd.Uc.GetXML("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	hd.Uc.DataInit(envelop)
}
