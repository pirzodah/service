package database

import (
	"avesto-service/internal/models"
	"context"
)

type Database interface {
	AddServices(ctx context.Context, service string) error
	GetServiceByID(ctx context.Context, service int64) (models.Services, error)
	UpdateService(ctx context.Context, service models.Services) error
	DeleteService(ctx context.Context, service int64) error

	AddInputs(ctx context.Context, inputs models.Inputs) error
	GetInputByID(ctx context.Context, inputs int64) (models.Inputs, error)
	UpdateInput(ctx context.Context, inputs models.Inputs) error
	DeleteInputs(ctx context.Context, inputs int64) error

	AddServiceInputs(ctx context.Context, serviceInput models.ServiceInputs) error
	GetServiceInputsByServiceID(ctx context.Context, serviceInputID int64) (models.ServiceInputs, error)
	UpdateServiceInput(ctx context.Context, serviceInput models.ServiceInputs) error
	DeleteServiceInput(ctx context.Context, serviceInputID int64) error
}
