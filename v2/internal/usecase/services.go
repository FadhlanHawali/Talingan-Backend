package usecase

import (
	"github.com/Talingan-Backend/v2/internal"
	"github.com/Talingan-Backend/v2/internal/entity"
)

type ServicesUsecase struct {
	dbRepo internal.ServicesDbRepo
}

func NewServicesUsecase(dbRepo internal.ServicesDbRepo) *ServicesUsecase {
	return &ServicesUsecase{
		dbRepo: dbRepo,
	}
}

func (s *ServicesUsecase) InsertNewKandang(kandang *entity.Kandang) (bool, error){
	flag,err := s.dbRepo.InsertNewKandang(kandang); if err != nil{
		return flag, err
	}
	return flag,nil
}
func (s *ServicesUsecase) GetKandangs()(bool,error,[]entity.Kandang){
	return s.dbRepo.GetKandangs()
}
func (s *ServicesUsecase) GetKandang(idKandang string) (bool, error,entity.Kandang){
	return s.dbRepo.GetKandang(idKandang)
}

func (s *ServicesUsecase) InsertMonitoring(monitoring *entity.DeteksiKandang) (bool,error){
	return s.dbRepo.InsertMonitoring(monitoring)
}
func (s *ServicesUsecase) GetMonitoring(idKandang string) (bool, error,[]entity.DeteksiKandang){
	return s.dbRepo.GetMonitoring(idKandang)
}

func (s *ServicesUsecase) GetNotification()(bool,error,[]entity.Notifikasi){
	return s.dbRepo.GetNotification()
}