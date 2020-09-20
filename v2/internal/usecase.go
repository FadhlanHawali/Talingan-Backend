package internal

import "github.com/Talingan-Backend/v2/internal/entity"

type ServicesUC interface{
	InsertNewKandang(kandang *entity.Kandang) (bool, error)
	GetKandangs()(bool,error,[]entity.Kandang)
	GetKandang(idKandang string) (bool, error,entity.Kandang)

	InsertMonitoring(monitoring *entity.DeteksiKandang) (bool,error)
	GetMonitoring(idKandang string) (bool, error,[]entity.DeteksiKandang)

	GetNotification()(bool,error,[]entity.Notifikasi)
}
