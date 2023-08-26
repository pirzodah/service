package entity

import (
	"avesto-service/internal/models"
	"context"
)

func (e *entity) AddInputs(ctx context.Context, inputs models.Inputs) error {
	return e.database.AddInputs(ctx, inputs)
}

func (e *entity) GetInputByID(ctx context.Context, params int64) (service models.Inputs, errCode error) {
	return e.database.GetInputByID(ctx, params)
}

func (e *entity) UpdateInput(ctx context.Context, params models.Inputs) error {
	return e.database.UpdateInput(ctx, params)
}

func (e *entity) DeleteInputs(ctx context.Context, params int64) error {
	return e.database.DeleteInputs(ctx, params)
}
