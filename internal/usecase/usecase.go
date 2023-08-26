package usecase

import (
	"avesto-service/internal/entity"
	"avesto-service/internal/models"
	"context"
)

type Usecase interface {
	AddServices(ctx context.Context, service models.Services) (errCode models.Response)
	GetServiceByID(ctx context.Context, serviceID int64) (service models.Services, errCode models.Response)
	UpdateService(ctx context.Context, service models.Services) (errCode models.Response)
	DeleteService(ctx context.Context, serviceID int64) (errCode models.Response)

	AddInputs(ctx context.Context, service models.Inputs) (errCode models.Response)
	GetInputs(ctx context.Context, serviceID int64) (service models.Inputs, errCode models.Response)
	UpdateInputs(ctx context.Context, service models.Inputs) (errCode models.Response)
	DeleteInputs(ctx context.Context, serviceID int64) (errCode models.Response)

	AddServiceInputs(ctx context.Context, serviceInput models.ServiceInputs) (errCode models.Response)
	GetServiceInputs(ctx context.Context, serviceInputID int64) (serviceInput models.ServiceInputs, errCode models.Response)
	UpdateServiceInputs(ctx context.Context, serviceInput models.ServiceInputs) (errCode models.Response)
	DeleteServiceInputs(ctx context.Context, serviceInputID int64) (errCode models.Response)
}

type usecase struct {
	entity entity.Entity
}

func New(entity entity.Entity) Usecase {
	return &usecase{
		entity: entity,
	}
}
