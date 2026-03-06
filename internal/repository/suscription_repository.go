package repository

import (
	"payment-service/internal/model"

	"gorm.io/gorm"
)

type SuscriptionRepository interface {
	Create(suscription *model.Suscription) error
	FindByID(id uint) (*model.Suscription, error)
	FindAll() ([]model.Suscription, error)
	Update(suscription *model.Suscription) error
	Delete(id uint) error
}

type suscriptionRepository struct {
	db *gorm.DB
}

func NewSuscriptionRepository(db *gorm.DB) SuscriptionRepository {
	return &suscriptionRepository{db: db}
}

func (r *suscriptionRepository) Create(data *model.Suscription) error {
	return r.db.Create(data).Error
}

func (r *suscriptionRepository) FindByID(id uint) (*model.Suscription, error) {
	var suscription model.Suscription
	err := r.db.First(&suscription, id).Error
	return &suscription, err
}

func (r *suscriptionRepository) FindAll() ([]model.Suscription, error) {
	var suscription []model.Suscription
	err := r.db.Find(&suscription).Error
	return suscription, err
}

func (r *suscriptionRepository) Update(suscription *model.Suscription) error {
	return r.db.Save(suscription).Error
}

func (r *suscriptionRepository) Delete(id uint) error {
	return r.db.Model(&model.Suscription{}).
		Where("id_suscription = ?", id).
		Update("status", "false").Error
}
