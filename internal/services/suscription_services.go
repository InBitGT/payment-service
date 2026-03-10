package services

import (
	"payment-service/internal/model"
	"payment-service/internal/repository"
)

type SuscriptionService interface {
	Create(data *model.Suscription) error
	GetAll() ([]model.Suscription, error)
	GetByID(id uint) (*model.Suscription, error)
	Delete(id uint) error
	Update(data *model.Suscription) error
}

type suscriptionService struct {
	repo repository.SuscriptionRepository
}

func NewSuscriptionService(repo repository.SuscriptionRepository) SuscriptionService {
	return &suscriptionService{repo}
}

func (s *suscriptionService) Create(data *model.Suscription) error {
	return s.repo.Create(data)
}

func (s *suscriptionService) GetAll() ([]model.Suscription, error) {
	return s.repo.FindAll()
}

func (s *suscriptionService) GetByID(id uint) (*model.Suscription, error) {
	return s.repo.FindByID(id)
}

func (s *suscriptionService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *suscriptionService) Update(data *model.Suscription) error {
	return s.repo.Update(data)
}
