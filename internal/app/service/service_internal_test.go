package service

import (
	"github.com/bemmanue/wildberries_L0/internal/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_HandleRoutes(t *testing.T) {
	s := newServer(teststore.New())

	testCases := []struct {
		name         string
		url          string
		expectedCode int
	}{
		{
			name:         "valid",
			url:          "/order",
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid",
			url:          "/orders",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, testCase.url, nil)
			s.ServeHTTP(rec, req)
			assert.Equal(t, testCase.expectedCode, rec.Code)
		})
	}

}
