package service

import (
	"errors"
	"testing"

	"github.com/pmaterer/meta/slip"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	CreateSlipFunc func(s slip.Slip) error
}

func (r *mockRepository) CreateSlip(s slip.Slip) error { return r.CreateSlipFunc(s) }

var (
	testSlip = slip.Slip{
		Body: "Lorem ipsum",
		Tags: []string{
			"tag1",
			"tag2",
			"tag3",
		},
	}
)

func TestCreateSlip(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(s slip.Slip) error
	}{
		{
			name:        "Create slip OK",
			errExpected: false,
			method: func(s slip.Slip) error {
				return nil
			},
		},
		{
			name:        "Create slip error",
			errExpected: true,
			method: func(s slip.Slip) error {
				return errors.New("oh no!")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mockRepository{CreateSlipFunc: tt.method}
			s := NewService(r)
			err := s.CreateSlip(testSlip)
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
