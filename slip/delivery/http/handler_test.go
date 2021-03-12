package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pmaterer/meta/slip"
	"github.com/stretchr/testify/assert"
)

var (
	testSlip = `{"body":"Lorem ipsum","tags":["tag1","tag2","tag3"]}`
)

type mockService struct {
	CreateSlipFunc func(s slip.Slip) error
}

func (s *mockService) CreateSlip(slip slip.Slip) error { return s.CreateSlipFunc(slip) }

func TestCreateSlip(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(slip.Slip) error
	}{
		{
			name:        "Create slip OK",
			errExpected: false,
			method: func(s slip.Slip) error {
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockService{
				CreateSlipFunc: tt.method,
			}
			h := NewHandler(s)

			r := gin.Default()
			r.POST("/slips", h.CreateSlip)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/slips", strings.NewReader(testSlip))
			r.ServeHTTP(w, req)

			if tt.errExpected {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
			} else {
				assert.Equal(t, http.StatusOK, w.Code)
			}

		})
	}
}
