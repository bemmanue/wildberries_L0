package server

import (
	"github.com/bemmanue/wildberries_L0/internal/cache/testcache"
	"github.com/bemmanue/wildberries_L0/internal/store/teststore"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_HandleRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := teststore.New()
	cache, _ := testcache.New(store)

	s := newServer(store, cache)

	testCases := []struct {
		name         string
		url          string
		expectedCode int
	}{
		{
			name:         "invalid",
			url:          "/orders",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, test.url, nil)
			s.ServeHTTP(rec, req)
			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}

}
