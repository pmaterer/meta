package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pmaterer/meta/slip"
	"github.com/stretchr/testify/assert"
)

var (
	testSlipPayload          = `{"body":"Lorem ipsum","tags":["tag1","tag2","tag3"]}`
	testSlipPayloadMalformed = `"body":"Lorem ipsum","tags":["tag1","tag2","tag3"]}`
	testSlipJSONResponse     = `{"id":1,"body":"Lorem ipsum","tags":["tag1","tag2","tag3"],"created_at":"2000-02-01T12:13:14.000000015Z","updated_at":"2000-02-01T12:13:14.000000015Z"}`
	testSlip                 = slip.Slip{
		ID:   1,
		Body: "Lorem ipsum",
		Tags: []string{
			"tag1",
			"tag2",
			"tag3",
		},
		CreatedAt: time.Date(2000, 2, 1, 12, 13, 14, 15, time.UTC),
		UpdatedAt: time.Date(2000, 2, 1, 12, 13, 14, 15, time.UTC),
	}
	testSlipsJSONResponse = `[{"id":2,"body":"Lorem ipsum","tags":["tag1","tag2","tag3"],"created_at":"2000-02-01T12:13:14.000000015Z","updated_at":"2000-02-01T12:13:14.000000015Z"},{"id":3,"body":"nothing to see here","tags":["a","b","c"],"created_at":"2000-02-01T12:13:14.000000015Z","updated_at":"2000-02-01T12:13:14.000000015Z"}]`
	testSlips             = []slip.Slip{
		{
			ID:   2,
			Body: "Lorem ipsum",
			Tags: []string{
				"tag1",
				"tag2",
				"tag3",
			},
			CreatedAt: time.Date(2000, 2, 1, 12, 13, 14, 15, time.UTC),
			UpdatedAt: time.Date(2000, 2, 1, 12, 13, 14, 15, time.UTC),
		},
		{
			ID:   3,
			Body: "nothing to see here",
			Tags: []string{
				"a",
				"b",
				"c",
			},
			CreatedAt: time.Date(2000, 2, 1, 12, 13, 14, 15, time.UTC),
			UpdatedAt: time.Date(2000, 2, 1, 12, 13, 14, 15, time.UTC),
		},
	}
)

type mockService struct {
	CreateSlipFunc  func(s slip.Slip) error
	GetSlipFunc     func(id int64) (slip.Slip, error)
	GetAllSlipsFunc func() ([]slip.Slip, error)
	UpdateSlipFunc  func(s slip.Slip) error
	DeleteSlipFunc  func(id int64) error
}

func (r *mockService) CreateSlip(s slip.Slip) error        { return r.CreateSlipFunc(s) }
func (r *mockService) GetSlip(id int64) (slip.Slip, error) { return r.GetSlipFunc(id) }
func (r *mockService) GetAllSlips() ([]slip.Slip, error)   { return r.GetAllSlipsFunc() }
func (r *mockService) UpdateSlip(s slip.Slip) error        { return r.UpdateSlipFunc(s) }
func (r *mockService) DeleteSlip(id int64) error           { return r.DeleteSlipFunc(id) }

func TestCreateSlip(t *testing.T) {
	tests := []struct {
		name             string
		errExpected      bool
		malformedPayload bool
		method           func(slip.Slip) error
	}{
		{
			name:             "Create slip OK",
			errExpected:      false,
			malformedPayload: false,
			method: func(s slip.Slip) error {
				return nil
			},
		},
		{
			name:             "Create slip malformed",
			errExpected:      true,
			malformedPayload: true,
			method: func(s slip.Slip) error {
				return nil
			},
		},
		{
			name:             "Create slip error",
			errExpected:      true,
			malformedPayload: false,
			method: func(s slip.Slip) error {
				return errors.New("kabam")
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
			if tt.malformedPayload {
				req, _ := http.NewRequest("POST", "/slips", strings.NewReader(testSlipPayloadMalformed))
				r.ServeHTTP(w, req)
			} else {
				req, _ := http.NewRequest("POST", "/slips", strings.NewReader(testSlipPayload))
				r.ServeHTTP(w, req)
			}

			if tt.errExpected {
				if tt.malformedPayload {
					assert.Equal(t, http.StatusBadRequest, w.Code)
				} else {
					assert.Equal(t, http.StatusInternalServerError, w.Code)
				}
			} else {
				assert.Equal(t, http.StatusOK, w.Code)
			}

		})
	}
}

func TestGetSlip(t *testing.T) {
	tests := []struct {
		name             string
		errExpected      bool
		malformedRequest bool
		method           func(id int64) (slip.Slip, error)
	}{
		{
			name:             "Get slip OK",
			errExpected:      false,
			malformedRequest: false,
			method: func(id int64) (slip.Slip, error) {
				return testSlip, nil
			},
		},
		{
			name:             "Get slip error",
			errExpected:      true,
			malformedRequest: false,
			method: func(id int64) (slip.Slip, error) {
				return testSlip, errors.New("bad times")
			},
		},
		{
			name:             "Get slip malformed",
			errExpected:      true,
			malformedRequest: true,
			method: func(id int64) (slip.Slip, error) {
				return testSlip, errors.New("bad times")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockService{
				GetSlipFunc: tt.method,
			}
			h := NewHandler(s)

			r := gin.Default()
			r.GET("/slips/:id", h.GetSlip)
			w := httptest.NewRecorder()

			if tt.malformedRequest {
				req, _ := http.NewRequest("GET", "/slips/x", nil)
				r.ServeHTTP(w, req)
			} else {
				req, _ := http.NewRequest("GET", "/slips/1", nil)
				r.ServeHTTP(w, req)
			}

			if tt.malformedRequest {
				assert.Equal(t, http.StatusBadRequest, w.Code)
			} else if tt.errExpected {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
			} else {
				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, testSlipJSONResponse, w.Body.String())
			}
		})
	}
}

func TestGetAllSlips(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func() ([]slip.Slip, error)
	}{
		{
			name:        "Get all slips OK",
			errExpected: false,
			method: func() ([]slip.Slip, error) {
				return testSlips, nil
			},
		},
		{
			name:        "Get all slips error",
			errExpected: true,
			method: func() ([]slip.Slip, error) {
				return testSlips, errors.New("uh oh")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockService{
				GetAllSlipsFunc: tt.method,
			}
			h := NewHandler(s)

			r := gin.Default()
			r.GET("/slips", h.GetAllSlips)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/slips", nil)
			r.ServeHTTP(w, req)

			if tt.errExpected {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
			} else {
				assert.Equal(t, http.StatusOK, w.Code)
				assert.Equal(t, testSlipsJSONResponse, w.Body.String())
			}
		})
	}
}

func TestUpdateSlip(t *testing.T) {
	tests := []struct {
		name             string
		errExpected      bool
		malformedPayload bool
		method           func(slip slip.Slip) error
	}{
		{
			name:             "Update slip OK",
			errExpected:      false,
			malformedPayload: false,
			method: func(s slip.Slip) error {
				return nil
			},
		},
		{
			name:             "Update slip error",
			errExpected:      true,
			malformedPayload: false,
			method: func(s slip.Slip) error {
				return errors.New("boom")
			},
		},
		{
			name:             "Update slip malformed payload",
			errExpected:      true,
			malformedPayload: true,
			method: func(s slip.Slip) error {
				return errors.New("boom")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockService{
				UpdateSlipFunc: tt.method,
			}
			h := NewHandler(s)

			r := gin.Default()
			r.PUT("/slips/:id", h.UpdateSlip)

			w := httptest.NewRecorder()

			if tt.malformedPayload {
				req, _ := http.NewRequest("PUT", "/slips/1", strings.NewReader(testSlipPayloadMalformed))
				r.ServeHTTP(w, req)
			} else {
				req, _ := http.NewRequest("PUT", "/slips/1", strings.NewReader(testSlipPayload))
				r.ServeHTTP(w, req)
			}

			if tt.errExpected {
				if tt.malformedPayload {
					assert.Equal(t, http.StatusBadRequest, w.Code)
				} else {
					assert.Equal(t, http.StatusInternalServerError, w.Code)
				}

			} else {
				assert.Equal(t, http.StatusOK, w.Code)
			}
		})
	}
}

func TestDeleteSlip(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(id int64) error
	}{
		{
			name:        "Delete slip OK",
			errExpected: false,
			method: func(id int64) error {
				return nil
			},
		},
		{
			name:        "Delete slip error",
			errExpected: true,
			method: func(id int64) error {
				return errors.New("boom")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &mockService{
				DeleteSlipFunc: tt.method,
			}
			h := NewHandler(s)

			r := gin.Default()
			r.DELETE("/slips/:id", h.DeleteSlip)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/slips/1", nil)
			r.ServeHTTP(w, req)

			if tt.errExpected {
				assert.Equal(t, http.StatusInternalServerError, w.Code)
			} else {
				assert.Equal(t, http.StatusOK, w.Code)
			}
		})
	}

}
