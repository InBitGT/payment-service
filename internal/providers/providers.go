package providers

import (
	"payment-service/internal/handlers"
	"payment-service/internal/repository"
	"payment-service/internal/services"

	"gorm.io/gorm"
)

type Provider struct {
	PlanHandler        *handlers.PlanHandler
	SuscriptionHandler *handlers.SuscriptionHandler
}

func NewProvider(db *gorm.DB) *Provider {
	//repo
	plan := repository.NewPlanRepository(db)
	suscription := repository.NewSuscriptionRepository(db)

	//services
	planservice := services.NewPlanService(plan)
	suscriptionservice := services.NewSuscriptionService(suscription)

	return &Provider{
		PlanHandler:        handlers.NewPlanHandler(planservice),
		SuscriptionHandler: handlers.NewSuscriptionHandler(suscriptionservice),
	}

}
