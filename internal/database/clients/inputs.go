package clients

import (
	"avesto-service/internal/models"
	"context"
)

func (db *db) AddInputs(ctx context.Context, input models.Inputs) error {
	_, err := db.postgres.Exec(ctx, "addInput", input.Name, input.TypeInputs)
	return err
}

func (db *db) GetInputByID(ctx context.Context, inputID int64) (models.Inputs, error) {
	var input models.Inputs
	err := db.postgres.QueryRow(ctx, "getInputByID", inputID).Scan(&input.ID, &input.Name, &input.TypeInputs)
	if err != nil {
		return input, err
	}
	return input, nil
}

func (db *db) UpdateInput(ctx context.Context, input models.Inputs) error {
	_, err := db.postgres.Exec(ctx, "updateInput", input.Name, input.TypeInputs, input.ID)
	return err
}

func (db *db) DeleteInputs(ctx context.Context, inputID int64) error {
	_, err := db.postgres.Exec(ctx, "deleteInput", inputID)
	return err
}
