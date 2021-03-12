package service

import "github.com/pmaterer/meta/slip"

type repository interface {
	CreateSlip(slip slip.Slip) error
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
