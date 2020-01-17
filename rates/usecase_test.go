package rates

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct {
	mock.Mock
}

func (m *UsecaseMock) GetRates(prefix string) (interface{}, error) {
	args := m.Called(prefix)
	return nil, args.Error(1)
}

func TestGetRatesCase1(t *testing.T) {
	t.Run("test get rates case lastes rate success", func(*testing.T) {
		t.Parallel()
		prefix := "/lastest-rate"
		mockRepo := new(RepositoryMock)
		mockRepo.On("GetLatestRate").Return(&LatestRate{}, nil)
		uc := NewUsecase(mockRepo)
		_, err := uc.GetRates(prefix)
		assert.Equal(t, nil, err)
	})
	t.Run("test get rates case lastes rate fail", func(t *testing.T) {
		t.Parallel()
		prefix := "/lastest-rate"
		mockRepo := new(RepositoryMock)
		mockRepo.On("GetLatestRate").Return(&LatestRate{}, errors.New("Error"))
		uc := NewUsecase(mockRepo)
		_, err := uc.GetRates(prefix)
		assert.NotEqual(t, nil, err)
	})
}

func TestGetRatesCase2(t *testing.T) {
	t.Run("test get rates case average currency success", func(*testing.T) {
		t.Parallel()
		prefix := "/average-currency"
		mockRepo := new(RepositoryMock)
		mockRepo.On("GetAverageCurrency").Return(&AverageRate{}, nil)
		uc := NewUsecase(mockRepo)
		_, err := uc.GetRates(prefix)
		assert.Equal(t, nil, err)
	})
	t.Run("test get rates case average currency fail", func(t *testing.T) {
		t.Parallel()
		prefix := "/average-currency"
		mockRepo := new(RepositoryMock)
		mockRepo.On("GetAverageCurrency").Return(&AverageRate{}, errors.New("Error"))
		uc := NewUsecase(mockRepo)
		_, err := uc.GetRates(prefix)
		assert.NotEqual(t, nil, err)
	})
}

func TestGetRatesCase3(t *testing.T) {
	t.Run("test get rates case default success", func(*testing.T) {
		t.Parallel()
		prefix := "/2020-01-01"
		prefix = strings.Replace(prefix, "/", "", -1)
		mockRepo := new(RepositoryMock)
		mockRepo.On("GetRateViaDate", prefix).Return(&RateViaDate{}, nil)
		uc := NewUsecase(mockRepo)
		_, err := uc.GetRates(prefix)
		assert.Equal(t, nil, err)
	})
	t.Run("test get rates case default fail", func(t *testing.T) {
		t.Parallel()
		prefix := "/2020-01-01"
		prefix = strings.Replace(prefix, "/", "", -1)
		mockRepo := new(RepositoryMock)
		mockRepo.On("GetRateViaDate", prefix).Return(&RateViaDate{}, errors.New("Error"))
		uc := NewUsecase(mockRepo)
		_, err := uc.GetRates(prefix)
		assert.NotEqual(t, nil, err)
	})
}
