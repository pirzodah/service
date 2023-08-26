package usecase

import (
	"avesto-service/internal/log"
	"avesto-service/internal/models"
	"context"
)

func (u *usecase) AddServiceInputs(ctx context.Context, serviceInput models.ServiceInputs) (errCode models.Response) {
	err := u.entity.AddServiceInputs(ctx, serviceInput)
	if err != nil {
		log.Log("usecase", "AddServiceInputs", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при добавлении данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}

func (u *usecase) GetServiceInputs(ctx context.Context, serviceInputID int64) (serviceInput models.ServiceInputs, errCode models.Response) {
	serviceInput, err := u.entity.GetServiceInputByID(ctx, serviceInputID)
	if err != nil {
		log.Log("usecase", "GetServiceInputs", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при получении данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}

func (u *usecase) UpdateServiceInputs(ctx context.Context, serviceInput models.ServiceInputs) (errCode models.Response) {
	err := u.entity.UpdateServiceInputs(ctx, serviceInput)
	if err != nil {
		log.Log("usecase", "UpdateServiceInputs", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при обновлении данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}

func (u *usecase) DeleteServiceInputs(ctx context.Context, serviceInputID int64) (errCode models.Response) {
	err := u.entity.DeleteServiceInputs(ctx, serviceInputID)
	if err != nil {
		log.Log("usecase", "DeleteServiceInputs", "ctrl.usecase.service.go", err)
		errCode.ErrorCode = models.INTERNALERROR
		errCode.Message = "Произошла ошибка при удалении данных"
		return
	}

	errCode.ErrorCode = models.SUCCESS
	return
}
