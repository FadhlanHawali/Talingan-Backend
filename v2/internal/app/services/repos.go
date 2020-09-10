package services

import (
	"github.com/Talingan-Backend/v2/internal"
	"gorm.io/gorm"
	db2 "github.com/Talingan-Backend/v2/internal/repo/db"
)

type Repos struct {
	servicesDBRepo internal.ServicesDbRepo
}


func newRepos(db *gorm.DB) (*Repos, error){
	r := &Repos{
		servicesDBRepo: db2.NewServicesDB(db),
	}

	return r,nil
}

func (r *Repos) Close() []error {
	var errs []error
	return errs
}
