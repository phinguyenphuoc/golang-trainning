package rates

import (
	"strings"
)

//UsecaseInterface interface
type UsecaseInterface interface {
	GetRates(prefix string) (interface{}, error)
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
