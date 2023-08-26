package entity

import (
	"avesto-service/internal/database"
	"avesto-service/internal/models"
	"context"
)

type Entity interface {
	AddServices(ctx context.Context, serviceName string) error
	GetServiceByID(ctx context.Context, serviceID int64) (service models.Services, errCode error)
	UpdateService(ctx context.Context, service models.Services) error
	DeleteService(ctx context.Context, serviceName int64) error

	AddInputs(ctx context.Context, service models.Inputs) error
	GetInputByID(ctx context.Context, service int64) (models.Inputs, error)
	UpdateInput(ctx context.Context, service models.Inputs) error
	DeleteInputs(ctx context.Context, service int64) error

	AddServiceInputs(ctx context.Context, serviceInput models.ServiceInputs) error
	GetServiceInputByID(ctx context.Context, serviceInputID int64) (models.ServiceInputs, error)
	UpdateServiceInputs(ctx context.Context, serviceInput models.ServiceInputs) error
	DeleteServiceInputs(ctx context.Context, serviceInputID int64) error
}

type entity struct {
	database database.Database
}

func New(database database.Database) Entity {
	return &entity{database: database}
}
