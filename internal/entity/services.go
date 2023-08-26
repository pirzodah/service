package entity

import (
	"avesto-service/internal/models"
	"context"
)

func (e *entity) AddServices(ctx context.Context, params string) error {
	return e.database.AddServices(ctx, params)
}

func (e *entity) GetServiceByID(ctx context.Context, params int64) (service models.Services, errCode error) {
	return e.database.GetServiceByID(ctx, params)
}

func (e *entity) UpdateService(ctx context.Context, params models.Services) error {
	return e.database.UpdateService(ctx, params)
}

func (e *entity) DeleteService(ctx context.Context, params int64) error {
	return e.database.DeleteService(ctx, params)
}
