package clients

import (
	"avesto-service/internal/models"
	"context"
)

func (db *db) AddServiceInputs(ctx context.Context, serviceInput models.ServiceInputs) error {
	tx, err := db.postgres.Begin(ctx)
	if err != nil {
		return err
	}
	for _, inputID := range serviceInput.InputIDs {
		_, err := tx.Exec(ctx, "addServiceInput",
			inputID,
			serviceInput.ServiceID,
			serviceInput.Required)
		if err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *db) GetServiceInputsByServiceID(ctx context.Context, serviceID int64) (models.ServiceInputs, error) {
	var serviceInputs models.ServiceInputs

	// Получение данных о service_inputs
	err := db.postgres.QueryRow(ctx, "getServiceInputsByServiceID", serviceID).Scan(
		&serviceInputs.ID, &serviceInputs.ServiceID, &serviceInputs.Required,
	)
	if err != nil {
		return serviceInputs, err
	}

	// Получение списка input_ids по service_id
	rows, err := db.postgres.Query(ctx, "getInputIDsByServiceID", serviceID)
	if err != nil {
		return serviceInputs, err
	}
	defer rows.Close()

	var inputIDs []int64
	for rows.Next() {
		var inputID int64
		err := rows.Scan(&inputID)
		if err != nil {
			return serviceInputs, err
		}
		inputIDs = append(inputIDs, inputID)
	}

	serviceInputs.InputIDs = inputIDs

	// Получение данных об инпутах по input_ids
	var inputs []models.Inputs
	for _, inputID := range inputIDs {
		var input models.Inputs
		err := db.postgres.QueryRow(ctx, "getInputByID", inputID).Scan(&input.ID, &input.Name, &input.TypeInputs)
		if err != nil {
			return serviceInputs, err
		}
		inputs = append(inputs, input)
	}

	serviceInputs.Inputs = inputs

	// Получение данных о сервисе
	err = db.postgres.QueryRow(ctx, "getServiceByID", serviceInputs.ServiceID).Scan(&serviceInputs.Services.ID, &serviceInputs.Services.Name)
	if err != nil {
		return serviceInputs, err
	}

	return serviceInputs, nil
}

func (db *db) UpdateServiceInput(ctx context.Context, serviceInput models.ServiceInputs) error {
	tx, err := db.postgres.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Close()

	// Обновляем основные данные о сервисе-вводе
	_, err = tx.Exec(ctx, "updateServiceInput",
		serviceInput.ServiceID,
		serviceInput.Required,
		serviceInput.ID)
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	// Обновляем связи между сервисом и вводами
	_, err = tx.Exec(ctx, "deleteServiceInputLinks", serviceInput.ID) // Удаляем существующие связи
	if err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	for _, inputID := range serviceInput.InputIDs {
		_, err := tx.Exec(ctx, "insertServiceInputLink",
			inputID,
			serviceInput.ID)
		if err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (db *db) DeleteServiceInput(ctx context.Context, serviceInputID int64) error {
	_, err := db.postgres.Exec(ctx, "deleteServiceInput", serviceInputID)
	return err
}
