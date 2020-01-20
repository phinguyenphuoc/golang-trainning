package rates

import (
	"errors"
	"exercise1/model"
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

func (m *UsecaseMock) GetXML(url string) model.Envelope {
	args := m.Called(url)
	return args.Get(0).(model.Envelope)
}

func (m *UsecaseMock) DataInit(data model.Envelope) error {
	args := m.Called(data)
	return args.Error(0)
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

func TestGetXML(t *testing.T) {
	mockRepo := new(RepositoryMock)
	uc := NewUsecase(mockRepo)
	data := uc.GetXML("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	assert.NotEmpty(t, data)
}

func TestDataInit(t *testing.T) {
	t.Run("test data init success", func(t *testing.T) {
		t.Parallel()
		item1 := model.CubeItem{Currency: "USD", Rate: "1.5"}
		item2 := model.CubeItem{Currency: "USD", Rate: "1.5"}
		list := []model.CubeItem{item1, item2}
		cube := model.Cube{Time: "2020-01-01", Cube: list}
		listCube := []model.Cube{cube, cube}
		cubemain := model.CubeMain{Cube: listCube}
		envelop := model.Envelope{Cube: cubemain}
		mockRepo := new(RepositoryMock)
		mockRepo.On("ImportData", "USD", "1.5", "2020-01-01").Return(nil)
		uc := NewUsecase(mockRepo)
		err := uc.DataInit(envelop)
		assert.Equal(t, nil, err)
	})
	t.Run("test data init fail", func(t *testing.T) {
		t.Parallel()
		item1 := model.CubeItem{Currency: "USD", Rate: "1.5"}
		item2 := model.CubeItem{Currency: "USD", Rate: "1.5"}
		list := []model.CubeItem{item1, item2}
		cube := model.Cube{Time: "2020-01-01", Cube: list}
		listCube := []model.Cube{cube, cube}
		cubemain := model.CubeMain{Cube: listCube}
		envelop := model.Envelope{Cube: cubemain}
		mockRepo := new(RepositoryMock)
		mockRepo.On("ImportData", "USD", "1.5", "2020-01-01").Return(errors.New("er"))
		uc := NewUsecase(mockRepo)
		err := uc.DataInit(envelop)
		assert.NotEqual(t, nil, err)
	})

}
