package service

import (
	"github.com/Talingan-Backend/v2/internal"
	"github.com/Talingan-Backend/v2/internal/app/services"
)

type Services struct {
	*ServicesService
}

func GetServices(app *services.ServicesRest) *Services{
	return &Services{
		ServicesService : NewServicesService(app),
	}
}

type ServicesService struct {
	uc internal.ServicesUC
}

func NewServicesService (app *services.ServicesRest) * ServicesService {
	return &ServicesService{
		uc: app.UseCases.Services,
	}
}

