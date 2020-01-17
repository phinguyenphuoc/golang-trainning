package rates

import (
	"errors"
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
