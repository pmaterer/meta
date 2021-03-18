package service

import "github.com/pmaterer/meta/slip"

type repository interface {
	CreateSlip(slip slip.Slip) error
	GetSlip(id int64) (slip.Slip, error)
	GetAllSlips() ([]slip.Slip, error)
	UpdateSlip(slip slip.Slip) error
	DeleteSlip(id int64) error
}

type Service struct {
	repository repository
}

func NewService(r repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) CreateSlip(slip slip.Slip) error {
	err := s.repository.CreateSlip(slip)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetSlip(id int64) (slip.Slip, error) {
	slip, err := s.repository.GetSlip(id)
	if err != nil {
		return slip, err
	}
	return slip, nil
}

func (s *Service) GetAllSlips() ([]slip.Slip, error) {
	slips, err := s.repository.GetAllSlips()
	if err != nil {
		return slips, err
	}
	return slips, nil
}

func (s *Service) UpdateSlip(slip slip.Slip) error {
	err := s.repository.UpdateSlip(slip)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteSlip(id int64) error {
	err := s.repository.DeleteSlip(id)
	if err != nil {
		return err
	}
	return nil
}
