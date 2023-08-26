package clients

import (
	"avesto-service/internal/models"
	"context"
)

func (db *db) AddServices(ctx context.Context, serviceName string) error {
	_, err := db.postgres.Exec(ctx, "addService", serviceName)
	return err
}

func (db *db) GetServiceByID(ctx context.Context, serviceID int64) (models.Services, error) {
	var service models.Services
	err := db.postgres.QueryRow(ctx, "getServiceByID", serviceID).Scan(&service.ID, &service.Name)
	if err != nil {
		return service, err
	}
	return service, nil
}

func (db *db) UpdateService(ctx context.Context, service models.Services) error {
	_, err := db.postgres.Exec(ctx, "updateService", service.Name, service.ID)
	return err
}

func (db *db) DeleteService(ctx context.Context, serviceID int64) error {
	_, err := db.postgres.Exec(ctx, "deleteService", serviceID)
	return err
}
