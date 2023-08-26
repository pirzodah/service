package controllers

import (
	"avesto-service/internal/models"
	"context"
)

type Controller interface {
	Serve(context.Context, *models.Configuration) error
	Shutdown(ctx context.Context) error
}
