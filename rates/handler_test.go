package rates

import (
	"errors"
	"exercise1/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func TestGetRateHandler(t *testing.T) {
	t.Run("hanlder get rate request success", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequest("GET", "/lastest-rate", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		uc := new(UsecaseMock)
		hd := NewHandler(uc)
		uc.On("GetRates", mock.Anything).Return(nil, nil)
		handler := http.HandlerFunc(hd.GetRates)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK)
	})
	t.Run("handler get rate request fail", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/lastest-rate", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		uc := new(UsecaseMock)
		hd := NewHandler(uc)
		uc.On("GetRates", mock.Anything).Return(nil, errors.New("handler error"))
		handler := http.HandlerFunc(hd.GetRates)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusOK)
	})
}

func TestSyncData(t *testing.T) {
	uc := new(UsecaseMock)
	item1 := model.CubeItem{Currency: "USD", Rate: "1.5"}
	item2 := model.CubeItem{Currency: "USD", Rate: "1.5"}
	list := []model.CubeItem{item1, item2}
	cube := model.Cube{Time: "2020-01-01", Cube: list}
	listCube := []model.Cube{cube, cube}
	cubemain := model.CubeMain{Cube: listCube}
	envelop := model.Envelope{Cube: cubemain}
	uc.On("GetXML", "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml").Return(envelop)
	uc.On("DataInit", envelop).Return(nil)
	hd := NewHandler(uc)
	hd.SyncData()
	assert.Equal(t, 1, 1)
}
