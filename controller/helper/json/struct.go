package json

import "time"

type ReturnData struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}
type Penjualan struct {
	ID               string    `json:"ID" bson:"ID"`
	NamaProduk       string    `json:"NamaProduk" bson:"NamaProduk"`
	JumlahPenjualan  string    `json:"JumlahPenjualan" bson:"JumlahPenjualan"`
	TanggalDataMasuk time.Time `json:"TanggalDataMasuk" bson:"TanggalDataMasuk"`
}
type Recap struct {
	ID             string `json:"ID" bson:"ID"`
	NamaPenjualan  string `json:"NamaPenjualan" bson:"NamaPenjualan"`
	RekapPenjualan string `json:"RekapPenjualan" bson:"RekapPenjualan"`
	JumlahKotor    int    `json:"JumlahKotor" bson:"JumlahKotor"`
	Pengeluaran    int    `json:"Pengeluaran" bson:"Pengeluaran"`
	JumlahBersih   int    `json:"JumlahBersih" bson:"JumlahBersih"`
}
type User struct {
	ID       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
type Pengeluaran struct {
	ID      string    `json:"ID" bson:"ID"`
	Nama    string    `json:"Nama" bson:"Nama"`
	Jumlah  int       `json:"Jumlah" bson:"Jumlah"`
	Tanggal time.Time `json:"Tanggal" bson:"Tanggal"`
}
