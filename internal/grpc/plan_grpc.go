package grpc

import (
	"context"
	"payment-service/internal/model"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	planpb "github.com/InBitGT/proto-definitions/payment/plan"
)

type PlanServer struct {
	planpb.UnimplementedPlanServiceServer
	db *gorm.DB
}

func NewPlanGRPCServer(db *gorm.DB) *PlanServer {
	return &PlanServer{db: db}
}

func (s *PlanServer) CreatePlan(ctx context.Context, req *planpb.CreatePlanRequest) (*planpb.PlanResponse, error) {
	plan := model.Plan{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
	}

	if err := s.db.WithContext(ctx).Create(&plan).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "error creating plan: %v", err)
	}

	return &planpb.PlanResponse{
		Id:          uint64(plan.ID),
		Name:        plan.Name,
		Price:       plan.Price,
		Description: plan.Description,
		Status:      plan.Status,
		CreatedAt:   plan.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   plan.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *PlanServer) GetPlan(ctx context.Context, req *planpb.GetPlanRequest) (*planpb.PlanResponse, error) {
	var plan model.Plan
	if err := s.db.WithContext(ctx).First(&plan, req.Id).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "plan %d not found", req.Id)
	}

	return &planpb.PlanResponse{
		Id:          uint64(plan.ID),
		Name:        plan.Name,
		Price:       plan.Price,
		Description: plan.Description,
		Status:      plan.Status,
		CreatedAt:   plan.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   plan.UpdatedAt.Format(time.RFC3339),
	}, nil
}
