package grpc

import (
	"context"
	"fmt"
	"payment-service/internal/model"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	suscriptionpb "github.com/InBitGT/proto-definitions/payment/suscription"
)

type SuscriptionServer struct {
	suscriptionpb.UnimplementedSuscriptionServiceServer
	db *gorm.DB
}

func NewSuscriptionGRPCServer(db *gorm.DB) *SuscriptionServer {
	return &SuscriptionServer{db: db}
}

func (s *SuscriptionServer) CreateSuscription(ctx context.Context, req *suscriptionpb.CreateSuscriptionRequest) (*suscriptionpb.SuscriptionResponse, error) {
	aStartedAt, err := time.Parse(time.RFC3339, req.AstartedAt)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid astarted_at: %v", err)
	}
	renewAt, err := time.Parse(time.RFC3339, req.RenewAt)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid renew_at: %v", err)
	}
	endAt, err := time.Parse(time.RFC3339, req.EndAt)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid end_at: %v", err)
	}

	sub := model.Suscription{
		TenantID:   uint(req.TenantId),
		PlanID:     uint(req.PlanId),
		AStartedAt: aStartedAt,
		RenewAt:    renewAt,
		EndAt:      endAt,
	}

	fmt.Println("valores de la sub", sub)

	if err := s.db.WithContext(ctx).Create(&sub).Error; err != nil {
		fmt.Println("valores de errpr", err)
		return nil, status.Errorf(codes.Internal, "error creating suscription: %v", err)
	}

	return toSuscriptionResponse(), nil
}

func (s *SuscriptionServer) GetSuscription(ctx context.Context, req *suscriptionpb.GetSuscriptionRequest) (*suscriptionpb.SuscriptionResponse, error) {
	var sub model.Suscription
	if err := s.db.WithContext(ctx).Preload("Plan").First(&sub, req.Id).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "suscription %d not found", req.Id)
	}

	return toSuscriptionResponse(), nil
}

func toSuscriptionResponse() *suscriptionpb.SuscriptionResponse {
	return &suscriptionpb.SuscriptionResponse{
		Success: true,
	}
}
