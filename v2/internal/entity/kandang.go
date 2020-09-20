package entity

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
	Accuracy float32 `json:"accuracy"`
}

type Notifikasi struct {
	ID int64 `gorm:"primary_key" json:"-"`
	Title string `json:"title"`
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}