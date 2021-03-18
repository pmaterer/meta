package service

import (
	"errors"
	"testing"

	"github.com/pmaterer/meta/slip"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	CreateSlipFunc  func(s slip.Slip) error
	GetSlipFunc     func(id int64) (slip.Slip, error)
	GetAllSlipsFunc func() ([]slip.Slip, error)
	UpdateSlipFunc  func(s slip.Slip) error
	DeleteSlipFunc  func(id int64) error
}

func (r *mockRepository) CreateSlip(s slip.Slip) error        { return r.CreateSlipFunc(s) }
func (r *mockRepository) GetSlip(id int64) (slip.Slip, error) { return r.GetSlipFunc(id) }
func (r *mockRepository) GetAllSlips() ([]slip.Slip, error)   { return r.GetAllSlipsFunc() }
func (r *mockRepository) UpdateSlip(s slip.Slip) error        { return r.UpdateSlipFunc(s) }
func (r *mockRepository) DeleteSlip(id int64) error           { return r.DeleteSlipFunc(id) }

var (
	testSlip = slip.Slip{
		ID:   1,
		Body: "Lorem ipsum",
		Tags: []string{
			"tag1",
			"tag2",
			"tag3",
		},
	}

	testSlips = []slip.Slip{
		{
			ID:   2,
			Body: "Lorem ipsum",
			Tags: []string{
				"tag1",
				"tag2",
				"tag3",
			},
		},
		{
			ID:   3,
			Body: "Fnord",
			Tags: []string{
				"all",
				"seeing",
				"eye",
			},
		},
		{
			ID:   4,
			Body: "There was a hole here. Now it's gone.",
			Tags: []string{
				"steam",
				"train",
			},
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
				return errors.New("oh no")
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

func TestGetSlip(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(id int64) (slip.Slip, error)
	}{
		{
			name:        "Get slip OK",
			errExpected: false,
			method: func(id int64) (slip.Slip, error) {
				return testSlip, nil
			},
		},
		{
			name:        "Get slip error",
			errExpected: true,
			method: func(id int64) (slip.Slip, error) {
				return testSlip, errors.New("oh no")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mockRepository{GetSlipFunc: tt.method}
			s := NewService(r)
			slip, err := s.GetSlip(1)
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testSlip.Body, slip.Body)
				assert.Equal(t, testSlip.Tags, slip.Tags)
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
				return testSlips, errors.New("this is bad")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mockRepository{GetAllSlipsFunc: tt.method}
			s := NewService(r)
			slips, err := s.GetAllSlips()
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
				for i, s := range slips {
					assert.Equal(t, testSlips[i].ID, s.ID)
					assert.Equal(t, testSlips[i].Body, s.Body)
					assert.Equal(t, testSlips[i].Tags, s.Tags)
				}
			}
		})
	}
}

func TestUpdateSlip(t *testing.T) {
	tests := []struct {
		name        string
		errExpected bool
		method      func(s slip.Slip) error
	}{
		{
			name:        "Update slip OK",
			errExpected: false,
			method: func(s slip.Slip) error {
				return nil
			},
		},
		{
			name:        "Update slip error",
			errExpected: true,
			method: func(s slip.Slip) error {
				return errors.New("things went wrong")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mockRepository{UpdateSlipFunc: tt.method}
			s := NewService(r)
			err := s.UpdateSlip(testSlip)
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
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
				return errors.New("kaboom")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mockRepository{DeleteSlipFunc: tt.method}
			s := NewService(r)
			err := s.DeleteSlip(1)
			if tt.errExpected {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
