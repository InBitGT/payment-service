package model

import "time"

type Plan struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;column:id_plan"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Price       string    `json:"price" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text"`
	Status      bool      `json:"status" gorm:"type:boolean;default:true"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt   time.Time `json:"update_at" gorm:"column:update_at;type:timestamptz;autoUpdateTime"`
}

func (Plan) TableName() string { return "payments.plan" }
