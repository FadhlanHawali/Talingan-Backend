package usecase

import "github.com/Talingan-Backend/v2/internal"

type ServicesUsecase struct {
	dbRepo internal.ServicesDbRepo
}

func NewServicesUsecase(dbRepo internal.ServicesDbRepo) *ServicesUsecase {
	return &ServicesUsecase{
		dbRepo: dbRepo,
	}
}
