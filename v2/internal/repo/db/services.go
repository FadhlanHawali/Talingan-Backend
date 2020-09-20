package db

import (
	"github.com/Talingan-Backend/v2/internal/entity"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
)

type ServicesDB struct {
	db *gorm.DB
}

func NewServicesDB(db *gorm.DB) *ServicesDB{
	return &ServicesDB{
		db: db,
	}
}

func (s *ServicesDB) InsertNewKandang(kandang *entity.Kandang) (bool, error){

	if err := s.db.Create(&kandang).Error;err!=nil{
		return false, errors.Errorf("invalid prepare statement :%+v\n", err)
	}
	return true, nil
}
func (s *ServicesDB) GetKandangs()(bool,error,[]entity.Kandang){
	var kandangs []entity.Kandang

	if err := s.db.Find(&kandangs).Error; err!=nil{
		return false,errors.Errorf("invalid prepare statement :%+v\n", err),[]entity.Kandang{}
	}else {
		return true,nil,kandangs
	}

}
func (s *ServicesDB) GetKandang(idKandang string) (bool, error,entity.Kandang){
	var kandang entity.Kandang
	if err := s.db.Where(&entity.Kandang{IdKandang:idKandang}).Find(&kandang).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return false,gorm.ErrRecordNotFound,entity.Kandang{}
		}else {
			return false,errors.Errorf("Error getting data"),entity.Kandang{}
		}
	}

	return true,nil,kandang
}

func (s *ServicesDB) InsertMonitoring(monitoring *entity.DeteksiKandang) (bool,error) {

	//if err := s.db.Where(&entity.DeteksiKandang{IdKandang: monitoring.IdKandang}).Find(&monitoring).Error; err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		if err = s.db.Create(&monitoring).Error;err!=nil{
	//			return false, errors.Errorf("invalid prepare statement :%+v\n", err)
	//		}
	//	} else {
	//		return false, errors.Errorf("Error inserting data")
	//	}
	//}

	return true,nil
}

func (s *ServicesDB) GetNotification()(bool,error,[]entity.Notifikasi){
	var notifications []entity.Notifikasi

	if err := s.db.Find(&notifications).Error; err!=nil{
		return false,errors.Errorf("invalid prepare statement :%+v\n", err),[]entity.Notifikasi{}
	}else {
		return true,nil,notifications
	}
}

func (s *ServicesDB) GetMonitoring(idKandang string) (bool, error,[]entity.DeteksiKandang){
	var monitor []entity.DeteksiKandang
	if err := s.db.
		Joins("LEFT JOIN hasil_deteksis on deteksi_kandangs.id = hasil_deteksis.id_deteksi_ayam_refer").
		Group("deteksi_kandangs.id").
		Preload("HasilDeteksi").
		Find(&monitor).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return false,gorm.ErrRecordNotFound,[]entity.DeteksiKandang{}
		}else {
			return false,errors.Errorf("Error getting data"),[]entity.DeteksiKandang{}
		}
	}

	return true,nil,monitor
}

func Paginate() func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi("1")
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi("7")
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}