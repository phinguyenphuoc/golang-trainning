package rates

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//HandlerInterface interface
type HandlerInterface interface {
	GetRates(w http.ResponseWriter, r *http.Request)
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
		fmt.Println("err")
	}
	json.NewEncoder(w).Encode(result)
}
