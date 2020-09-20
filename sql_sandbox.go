
package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)





var db *gorm.DB
type Kandang struct {
	ID int64 `gorm:"primary_key" json:"-"`
	IdKandang string `json:"id_kandang"`
	NamaKandang string `json:"nama_kandang"`
	NamaTiang string `json:"nama_tiang"`
	JumlahAyam int `json:"jumlah_ayam"`
	JenisAyam string `json:"jenis_ayam"`
	Usia int `json:"usia"`
	Suhu int `json:"suhu"`
	DeteksiKandang []DeteksiKandang `gorm:"ForeignKey:IdKandangRefer" json:"deteksi_kandang,omitempty"`
}


type DeteksiKandang struct {
	ID int64 `gorm:"primary_key" json:"-"`
	IdKandangRefer int64 `json:"-"`
	IdDeteksi string `json:"id_deteksi"`
	TimestampStart int64 `json:"timestamp_start"`
	TimestampEnd int64 `json:"timestamp_end"`
	HasilDeteksi []HasilDeteksi `gorm:"ForeignKey:IdDeteksiAyamRefer" json:"hasil_deteksi"`
}

type HasilDeteksi struct {
	ID int64 `gorm:"primary_key" json:"-"`
	IdHasilDeteksi string `json:"id_hasil_deteksi"`
	IdDeteksiAyamRefer int64 `json:"-"`
	Timestamp int64 `json:"timestamp"`
	FrekuensiDengkuran int64 `json:"frekuensi_dengkuran"`
	Accuracy float64 `json:"accuracy"`
}

type Notifikasi struct {
	ID int64 `gorm:"primary_key" json:"-"`
	Title string `json:"title"`
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}
func main() {
	var err error
	dsn := "talingan:RahasiaTalinganku2020%@/talingan?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.Logger.LogMode(0)
	db.Migrator().DropTable(HasilDeteksi{},DeteksiKandang{},Kandang{},Notifikasi{})
	db.AutoMigrate(new (Kandang),new(DeteksiKandang),new(HasilDeteksi),new(Notifikasi))

	addDeteksiAyam()
}

func addDeteksiAyam(){
	var kandang Kandang
	var deteksiKandang []DeteksiKandang
	var hasilDeteksi []HasilDeteksi
	var notifikasi []Notifikasi
	var OneHour,NextDay int64
	//max :=8
	min := 0
	OneHour = 0
	NextDay = 0;
	jmlPink := 0;
	for i:=0;i<7;i++ {
		for j:=0;j<24;j++ {
			frekDengkuran := randomInt(min,2)
			if frekDengkuran == 1{
				jmlPink++
			}
			if jmlPink >= 90{
				frekDengkuran = 0;
			}
			if j == 21 && i ==1{
				frekDengkuran = 5
			}else if j ==19 && i == 3{
				frekDengkuran = 7
			}else if j == 10 && i ==6{
				frekDengkuran = 4
			}
			hasilDeteksi = append(hasilDeteksi,HasilDeteksi{
				IdHasilDeteksi:     fmt.Sprintf("%d-dtksi-%d",i,j),
				Timestamp:     1599955200+OneHour,
				FrekuensiDengkuran: frekDengkuran,
				Accuracy: randomFloat(60,70),
			})
			if frekDengkuran >4 {
				notifikasi = append(notifikasi,Notifikasi{
					Title:     fmt.Sprintf("Ada Ayam Mendengkur Sebanyak %d!",frekDengkuran),
					Message:   fmt.Sprintf("Segera cek Kandang Barat Tiang 1"),
					Timestamp: 1599955200+OneHour,
				})
			}
			OneHour += 3600
		}
		deteksiKandang = append(deteksiKandang, DeteksiKandang{
			IdDeteksi:      fmt.Sprintf("id-dtks-%d",i),
			TimestampStart: 1599955200+NextDay,
			TimestampEnd:   1600041599+NextDay,
			HasilDeteksi: hasilDeteksi,
		},
		)
		NextDay += 86400
		log.Println(NextDay)
		hasilDeteksi = []HasilDeteksi{}
	}

	kandang = Kandang{
		IdKandang:      "id-kdg-1",
		NamaKandang:    "Kandang Barat",
		NamaTiang:      "Tiang 1",
		JumlahAyam:     100,
		JenisAyam:      "Broiler",
		Usia:           10,
		Suhu:           10,
		DeteksiKandang: deteksiKandang,
	}

	err := db.Create(&kandang);if err != nil {
		log.Println(err.Error)
	}

	err = db.Create(&notifikasi); if err !=nil{
		log.Println(err.Error)
	}
	log.Println("success")
}

func randomInt(min int, max int) int64{
	return int64(rand.Intn(max - min) + min)
}

func randomFloat(min ,max float64) float64{
	r := min + rand.Float64() * (max - min)
	return math.Floor(r*100)/100
}
