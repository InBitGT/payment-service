package model

import "time"

type Suscription struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement;column:id_suscription"`
	TenantID   uint      `json:"tenant_id" gorm:"column:tenant_id;not null;index"`
	PlanID     uint      `json:"plan_id" gorm:"column:plan_id;not null;index"`
	Plan       *Plan     `json:"plan" gorm:"foreignKey:PlanID;references:ID"`
	AStartedAt time.Time `json:"astarted_at" gorm:"column:astarted_at;type:timestamptz"`
	RenewAt    time.Time `json:"renew_at" gorm:"column:renew_at;type:timestamptz"`
	EndAt      time.Time `json:"end_at" gorm:"column:end_at;type:timestamptz"`
	Status     bool      `json:"status" gorm:"type:boolean;default:true"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt  time.Time `json:"update_at" gorm:"column:update_at;type:timestamptz;autoUpdateTime"`
}

func (Suscription) TableName() string { return "suscription" }
