package json

import "time"

type ReturnData struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}
type Penjualan struct {
	ID               string    `json:"id" bson:"id"`
	NamaProduk       string    `json:"namaproduk" bson:"namaproduk"`
	JumlahPenjualan  string    `json:"jumlahpenjualan" bson:"jumlahpenjualan"`
	TanggalDataMasuk time.Time `json:"tanggaldatamasuk" bson:"tanggaldatamasuk"`
}
type Recap struct {
	ID             string `json:"id" bson:"id"`
	NamaPenjualan  string `json:"namapenjualan" bson:"namapenjualan"`
	RekapPenjualan string `json:"rekappenjualan" bson:"rekappenjualan"`
	JumlahKotor    int    `json:"jumlahkotor" bson:"jumlahkotor"`
	Pengeluaran    int    `json:"pengeluaran" bson:"pengeluaran"`
	JumlahBersih   int    `json:"jumlahbersih" bson:"jumlahbersih"`
}
type User struct {
	ID       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
type Pengeluaran struct {
	ID      string    `json:"ID" bson:"ID"`
	Nama    string    `json:"nama" bson:"nama"`
	Jumlah  int       `json:"jumlah" bson:"jumlah"`
	Tanggal time.Time `json:"tanggal" bson:"tanggal"`
}
