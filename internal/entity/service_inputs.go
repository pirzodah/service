package entity

import (
	"avesto-service/internal/models"
	"context"
)

func (e *entity) AddServiceInputs(ctx context.Context, serviceInput models.ServiceInputs) error {
	return e.database.AddServiceInputs(ctx, serviceInput)
}

func (e *entity) GetServiceInputByID(ctx context.Context, serviceInputID int64) (models.ServiceInputs, error) {
	return e.database.GetServiceInputsByServiceID(ctx, serviceInputID)
}

func (e *entity) UpdateServiceInputs(ctx context.Context, serviceInput models.ServiceInputs) error {
	return e.database.UpdateServiceInput(ctx, serviceInput)
}

func (e *entity) DeleteServiceInputs(ctx context.Context, serviceInputID int64) error {
	return e.database.DeleteServiceInput(ctx, serviceInputID)
}
