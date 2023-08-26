package usecase

import (
	"avesto-service/internal/log"
	"avesto-service/internal/models"
	"context"
)

func (u *usecase) AddInputs(ctx context.Context, service models.Inputs) (errCode models.Response) {

	if len(service.Name) <= 0 {
		log.Log("usecase", "len(service.Name)", "ctrl.usecase.service.go", nil)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при добавление данных"
		return
	}
	err := u.entity.AddInputs(ctx, service)
	if err != nil {
		log.Log("usecase", "AddServices", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при добавление данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}

func (u *usecase) GetInputs(ctx context.Context, serviceID int64) (service models.Inputs, errCode models.Response) {

	service, err := u.entity.GetInputByID(ctx, serviceID)
	if err != nil {
		log.Log("usecase", "GetServiceByID", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при добавление данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}

func (u *usecase) UpdateInputs(ctx context.Context, service models.Inputs) (errCode models.Response) {

	if len(service.Name) <= 0 {
		log.Log("usecase", "len(service.Name)", "ctrl.usecase.service.go", nil)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при добавление данных"
		return
	}
	err := u.entity.UpdateInput(ctx, service)
	if err != nil {
		log.Log("usecase", "AddServices", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при добавление данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}

func (u *usecase) DeleteInputs(ctx context.Context, serviceID int64) (errCode models.Response) {

	err := u.entity.DeleteInputs(ctx, serviceID)
	if err != nil {
		log.Log("usecase", "AddServices", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при добавление данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}
