package routes

import (
	"payment-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterSuscriptionRoutes(r *gin.RouterGroup, h *handlers.SuscriptionHandler) {
	roles := r.Group("/suscription")
	{
		roles.POST("", h.Create)
		roles.GET("", h.GetAll)
		roles.GET("/:id", h.GetByID)
		roles.PUT("/:id", h.Update)
		roles.DELETE("/:id", h.Delete)
	}
}
