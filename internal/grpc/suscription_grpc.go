package grpc

import (
	"context"
	"payment-service/internal/model"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	planpb "github.com/InBitGT/proto-definitions/payment/plan"
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

	if err := s.db.WithContext(ctx).Create(&sub).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "error creating suscription: %v", err)
	}

	return toSuscriptionResponse(&sub), nil
}

func (s *SuscriptionServer) GetSuscription(ctx context.Context, req *suscriptionpb.GetSuscriptionRequest) (*suscriptionpb.SuscriptionResponse, error) {
	var sub model.Suscription
	if err := s.db.WithContext(ctx).Preload("Plan").First(&sub, req.Id).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "suscription %d not found", req.Id)
	}

	return toSuscriptionResponse(&sub), nil
}

func toSuscriptionResponse(sub *model.Suscription) *suscriptionpb.SuscriptionResponse {
	return &suscriptionpb.SuscriptionResponse{
		Id:       uint64(sub.ID),
		TenantId: uint64(sub.TenantID),
		PlanId:   uint64(sub.PlanID),
		Plan: &planpb.PlanResponse{
			Id:          uint64(sub.Plan.ID),
			Name:        sub.Plan.Name,
			Price:       sub.Plan.Price,
			Description: sub.Plan.Description,
			Status:      sub.Plan.Status,
			CreatedAt:   sub.Plan.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   sub.Plan.UpdatedAt.Format(time.RFC3339),
		},
		AstartedAt: sub.AStartedAt.Format(time.RFC3339),
		RenewAt:    sub.RenewAt.Format(time.RFC3339),
		EndAt:      sub.EndAt.Format(time.RFC3339),
		Status:     sub.Status,
		CreatedAt:  sub.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  sub.UpdatedAt.Format(time.RFC3339),
	}
}
