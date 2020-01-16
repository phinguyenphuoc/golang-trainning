package rates

import (
	"database/sql"
	"fmt"
)

//RepositoryInterface interface
type RepositoryInterface interface {
	GetLatestRate() (*LatestRate, error)
	GetRateViaDate(date string) (*RateViaDate, error)
	GetAverageCurrency() (*AverageRate, error)
	GetLastestDate() (string, error)
}

//Repository struct
type Repository struct {
	DB *sql.DB
}

//NewRepository func
func NewRepository(dbi *sql.DB) RepositoryInterface {
	return &Repository{dbi}
}

//GetLatestRate func
func (rp *Repository) GetLatestRate() (*LatestRate, error) {
	var rateResult LatestRate
	data := make(map[string]string)
	lastestDate, err := rp.GetLastestDate()
	// fmt.Println(lastestDate, err)
	if err != nil {
		return &LatestRate{}, nil
	}
	//get data from db
	query := fmt.Sprintf("SELECT currency, rate FROM Cube WHERE reg_date LIKE '%s'", lastestDate)
	results, err := rp.DB.Query(query)

	if err != nil {
		return &LatestRate{}, nil
	}

	//loop through data
	for results.Next() {
		var currency string
		var rate string
		err = results.Scan(&currency, &rate)
		data[currency] = rate
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	rateResult.Rate = data
	return &rateResult, nil
}

//GetRateViaDate func
func (rp *Repository) GetRateViaDate(date string) (*RateViaDate, error) {
	var rateViaDate RateViaDate
	data := make(map[string]string)

	query := fmt.Sprintf("SELECT currency, rate FROM Cube WHERE reg_date LIKE '%s'", date)
	result, err := rp.DB.Query(query)
	if err != nil {
		return &RateViaDate{}, err
	}
	for result.Next() {
		var cur string
		var rate string
		result.Scan(&cur, &rate)
		data[cur] = rate
	}
	rateViaDate.Date = date
	rateViaDate.Rate = data
	return &rateViaDate, nil

}

//GetAverageCurrency function
func (rp *Repository) GetAverageCurrency() (*AverageRate, error) {
	var averageRate AverageRate
	data := make(map[string]Average)
	result, err := rp.DB.Query("SELECT currency, min(rate), max(rate), avg(rate) FROM Cube group by currency")
	if err != nil {
		return &AverageRate{}, err
	}
	for result.Next() {
		var cur string
		var min string
		var max string
		var avg string
		result.Scan(&cur, &min, &max, &avg)
		data[cur] = Average{min, max, avg}
	}
	averageRate.Rate = data
	return &averageRate, nil
}

//GetLastestDate function
func (rp *Repository) GetLastestDate() (string, error) {
	var date string
	result, err := rp.DB.Query("Select reg_date from Cube order by reg_date desc limit 1")
	if err != nil {
		return "2000-01-01", err
	}
	for result.Next() {
		result.Scan(&date)
	}

	fmt.Println(date)
	return date, nil
}