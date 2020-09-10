package services

import (
	"github.com/Talingan-Backend/v2/internal"
	"github.com/Talingan-Backend/v2/internal/usecase"
)

type Usecases struct {
	Services internal.ServicesUC
}

func newUsecases(repos *Repos) *Usecases {
	return &Usecases{
		Services: usecase.NewServicesUsecase(repos.servicesDBRepo),
	}
}

func (*Usecases) Close() []error{
	var errs []error
	return errs
}
