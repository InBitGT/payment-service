package server

import (
	"log"
	"net"
	"payment-service/internal/handlers"
	"payment-service/internal/repository"
	"payment-service/internal/routes"
	"payment-service/internal/services"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type App struct {
	grpcServer *grpc.Server
	router     *gin.Engine
}

func NewApp(db *gorm.DB) *App {
	// Inicializar dependencias
	plan := repository.NewPlanRepository(db)
	planservice := services.NewPlanService(plan)
	planHandler := handlers.NewPlanHandler(planservice)

	suscription := repository.NewSuscriptionRepository(db)
	suscriptionservice := services.NewSuscriptionService(suscription)
	suscriptionHandler := handlers.NewSuscriptionHandler(suscriptionservice)

	// Router Gin
	router := gin.Default()
	api := router.Group("/payment-service/api")
	routes.RegisterPlanRoutes(api, planHandler)
	routes.RegisterSuscriptionRoutes(api, suscriptionHandler)

	// Servidor gRPC
	s := grpc.NewServer()

	return &App{grpcServer: s, router: router}
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
