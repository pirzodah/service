package usecase

import (
	"avesto-service/internal/log"
	"avesto-service/internal/models"
	"context"
)

func (u *usecase) AddServices(ctx context.Context, service models.Services) (errCode models.Response) {

	if len(service.Name) <= 0 {
		log.Log("usecase", "len(service.Name)", "ctrl.usecase.service.go", nil)
		errCode.ErrorCode = models.INTERNALERROR
		return
	}

	err := u.entity.AddServices(ctx, service.Name)
	if err != nil {
		log.Log("usecase", "AddServices", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при добавление данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}

func (u *usecase) GetServiceByID(ctx context.Context, serviceID int64) (service models.Services, errCode models.Response) {

	service, err := u.entity.GetServiceByID(ctx, serviceID)
	if err != nil {
		log.Log("usecase", "GetServiceByID", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при получение данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}

func (u *usecase) UpdateService(ctx context.Context, service models.Services) (errCode models.Response) {

	if len(service.Name) <= 0 {
		log.Log("usecase", "len(service.Name)", "ctrl.usecase.service.go", nil)
		errCode.ErrorCode = models.INTERNALERROR
		return
	}
	err := u.entity.UpdateService(ctx, service)
	if err != nil {
		log.Log("usecase", "UpdateService", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при обновление данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}

func (u *usecase) DeleteService(ctx context.Context, serviceID int64) (errCode models.Response) {

	err := u.entity.DeleteService(ctx, serviceID)
	if err != nil {
		log.Log("usecase", "DeleteService", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при удаление данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}
