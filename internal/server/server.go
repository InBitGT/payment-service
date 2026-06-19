package server

import (
	"log"
	"net"
	"payment-service/internal/providers"
	"payment-service/internal/routes"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type App struct {
	grpcServer *grpc.Server
	router     *gin.Engine
}

func NewApp(db *gorm.DB) *App {
	c := providers.NewProvider(db)

	router := gin.Default()
	api := router.Group("/payment-service/api")
	routes.RegisterPlanRoutes(api, c.PlanHandler)
	routes.RegisterSuscriptionRoutes(api, c.SuscriptionHandler)

	return &App{
		grpcServer: grpc.NewServer(),
		router:     router,
	}
}

func (a *App) Run(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error al escuchar:", err)
	}
	log.Println("User Service corriendo en", port)
	if err := a.grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func (a *App) RunHTTP(port string) {
	log.Println("HTTP corriendo en :", port)
	if err := a.router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
