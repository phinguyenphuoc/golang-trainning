package rates

import (
	"encoding/xml"
	"exercise1/model"
	"io/ioutil"
	"net/http"
	"strings"
)

//UsecaseInterface interface
type UsecaseInterface interface {
	GetRates(prefix string) (interface{}, error)
	GetXML(url string) model.Envelope
	DataInit(data model.Envelope) error
}

//Usecase struct
type Usecase struct {
	Repo RepositoryInterface
}

//NewUsecase funciton
func NewUsecase(rp RepositoryInterface) UsecaseInterface {
	return &Usecase{rp}
}

//GetRates function
func (uc *Usecase) GetRates(prefix string) (interface{}, error) {
	switch prefix {
	case "/lastest-rate":
		result, err := uc.Repo.GetLatestRate()
		if err != nil {
			return nil, err
		}
		return result, nil
	case "/average-currency":
		result, err := uc.Repo.GetAverageCurrency()
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		prefix = strings.Replace(prefix, "/", "", -1)
		result, err := uc.Repo.GetRateViaDate(prefix)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}

//DataInit function
func (uc *Usecase) DataInit(data model.Envelope) error {
	for i, v := range data.Cube.Cube {
		for _, value := range data.Cube.Cube[i].Cube {
			err := uc.Repo.ImportData(value.Currency, value.Rate, v.Time)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//GetXML function
func (uc *Usecase) GetXML(url string) model.Envelope {
	resp, _ := http.Get(url)
	rawXMLData, _ := ioutil.ReadAll(resp.Body)
	var data model.Envelope
	xml.Unmarshal([]byte(rawXMLData), &data)
	return data
}
