package services

import (
	"payment-service/internal/model"
	"payment-service/internal/repository"
)

type PlanService interface {
	Create(data *model.Plan) error
	GetAll() ([]model.Plan, error)
	GetByID(id uint) (*model.Plan, error)
	Delete(id uint) error
	Update(data *model.Plan) error
}

type planService struct {
	repo repository.PlanRepository
}

func NewPlanService(repo repository.PlanRepository) PlanService {
	return &planService{repo}
}

func (s *planService) Create(data *model.Plan) error {
	return s.repo.Create(data)
}

func (s *planService) GetAll() ([]model.Plan, error) {
	return s.repo.FindAll()
}

func (s *planService) GetByID(id uint) (*model.Plan, error) {
	return s.repo.FindByID(id)
}

func (s *planService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *planService) Update(data *model.Plan) error {
	return s.repo.Update(data)
}
