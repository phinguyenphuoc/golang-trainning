package rates

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
)

//TestGet func
type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetLatestRate() (*LatestRate, error) {
	args := m.Called()
	return args.Get(0).(*LatestRate), args.Error(1)
}

func (m *RepositoryMock) GetRateViaDate(date string) (*RateViaDate, error) {
	args := m.Called(date)
	return args.Get(0).(*RateViaDate), args.Error(1)
}

func (m *RepositoryMock) GetAverageCurrency() (*AverageRate, error) {
	args := m.Called()
	return args.Get(0).(*AverageRate), args.Error(0)
}

func (m *RepositoryMock) GetLastestDate() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(0)
}

func TestGetLatestRate(t *testing.T) {
	t.Run("test get latest rate success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		rows := sqlmock.NewRows([]string{"date"}).
			AddRow("2020-01-01")
		mock.ExpectQuery("Select reg_date from Cube order by reg_date desc limit 1").
			WillReturnRows(rows)
		rows2 := sqlmock.NewRows([]string{"currency", "rate"}).
			AddRow("USD", "1.2432").
			AddRow("VND", "2.562")
		mock.ExpectQuery("SELECT currency, rate FROM Cube WHERE reg_date LIKE '2020-01-01'").
			WillReturnRows(rows2)
		rp := NewRepository(db)
		result, err := rp.GetLatestRate()
		assert.Equal(t, nil, err)
		assert.NotEqual(t, nil, result)
	})
	t.Run("test get lastest rate fail since get lastest date fail", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectQuery("Select reg_date from Cube order by reg_date desc limit 1").
			WillReturnError(errors.New("DB error"))
		rp := NewRepository(db)
		result, err := rp.GetLatestRate()
		assert.NotEqual(t, nil, err)
		assert.Empty(t, result)
	})
	t.Run("test get lastest rate fail ", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		rows := sqlmock.NewRows([]string{"date"}).
			AddRow("2020-01-01")
		mock.ExpectQuery("Select reg_date from Cube order by reg_date desc limit 1").
			WillReturnRows(rows)
		mock.ExpectQuery("SELECT currency, rate FROM Cube WHERE reg_date LIKE '2020-01-01'").
			WillReturnError(errors.New("DB error"))
		rp := NewRepository(db)
		result, err := rp.GetLatestRate()
		assert.NotEqual(t, nil, err)
		assert.Empty(t, result)
	})
}

func TestGetRateViaDate(t *testing.T) {
	t.Run("get rate via date success", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		rows := sqlmock.NewRows([]string{"currency", "rate"}).
			AddRow("USD", "1.2432")
		mock.ExpectQuery("SELECT currency, rate FROM Cube WHERE reg_date LIKE '2020-01-01'").
			WillReturnRows(rows)
		rp := NewRepository(db)
		result, err := rp.GetRateViaDate("2020-01-01")
		assert.NotEqual(t, nil, result)
		assert.Equal(t, nil, err)
	})
	t.Run("get rate via date fail", func(t *testing.T) {
		t.Parallel()
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectQuery("SELECT currency, rate FROM Cube WHERE reg_date LIKE '2020-01-01'").
			WillReturnError(errors.New("DB error"))
		rp := NewRepository(db)
		result, err := rp.GetRateViaDate("2020-01-01")
		assert.Empty(t, result)
		assert.NotEqual(t, nil, err)
	})
}

func TestGetAverageCurrency(t *testing.T) {
	t.Run("test get average currency success", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		rows := sqlmock.NewRows([]string{"currency", "min\\(rate\\)", "max\\(rate\\)", "avg\\(rate\\)"}).
			AddRow("USD", "1.2432", "2.5232", "1.6887").
			AddRow("VND", "1.1232", "2.053", "1.5212")
		mock.ExpectQuery(`SELECT currency, min\(rate\), max\(rate\), avg\(rate\) FROM Cube group by currency`).
			WillReturnRows(rows)
		rp := NewRepository(db)
		result, err := rp.GetAverageCurrency()
		assert.NotEqual(t, nil, result)
		assert.Equal(t, nil, err)
	})
	t.Run("test get average currency fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectQuery(`SELECT currency, min\(rate\), max\(rate\), avg\(rate\) FROM Cube group by currency`).
			WillReturnError(errors.New("DB error"))
		rp := NewRepository(db)
		result, err := rp.GetAverageCurrency()
		assert.Empty(t, result)
		assert.NotEqual(t, nil, err)
	})
}
