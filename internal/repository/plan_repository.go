package repository

import (
	"payment-service/internal/model"

	"gorm.io/gorm"
)

type PlanRepository interface {
	Create(plan *model.Plan) error
	FindByID(id uint) (*model.Plan, error)
	FindAll() ([]model.Plan, error)
	Update(plan *model.Plan) error
	Delete(id uint) error
}

type planRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) PlanRepository {
	return &planRepository{db: db}
}

func (r *planRepository) Create(data *model.Plan) error {
	return r.db.Create(data).Error
}

func (r *planRepository) FindByID(id uint) (*model.Plan, error) {
	var plan model.Plan
	err := r.db.Where("status = ?", true).First(&plan, id).Error
	return &plan, err
}

func (r *planRepository) FindAll() ([]model.Plan, error) {
	var plan []model.Plan
	err := r.db.Where("status = ?", true).Find(&plan).Error
	return plan, err
}

func (r *planRepository) Update(plan *model.Plan) error {
	return r.db.Model(&model.Plan{}).
		Where("id_plan = ? AND status = ?", plan.ID, true).
		Updates(plan).Error
}

func (r *planRepository) Delete(id uint) error {
	return r.db.Model(&model.Plan{}).
		Where("id_plan = ?", id).
		Update("status", false).Error
}
