package config

import (
	"fmt"
	"log"
	"net"
	grpcServer "payment-service/internal/grpc"

	planpb "github.com/InBitGT/proto-definitions/payment/plan"
	suscriptionpb "github.com/InBitGT/proto-definitions/payment/suscription"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func StartGRPCServer(db *gorm.DB, port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("gRPC listen error: %v", err)
	}

	srv := grpc.NewServer()
	planpb.RegisterPlanServiceServer(srv, grpcServer.NewPlanGRPCServer(db))
	suscriptionpb.RegisterSuscriptionServiceServer(srv, grpcServer.NewSuscriptionGRPCServer(db))

	log.Printf("gRPC server listening on :%s", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("gRPC serve error: %v", err)
	}
}
