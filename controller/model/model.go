package model

type ReturnData struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}
type Penjualan struct {
	ID               int    `json:"id" bson:"id"`
	NamaProduk       string `json:"namaproduk" bson:"namaproduk"`
	JumlahPenjualan  int    `json:"jumlahpenjualan" bson:"jumlahpenjualan"`
	TanggalDataMasuk string `json:"tanggaldatamasuk" bson:"tanggaldatamasuk"`
	Cabang           string `json:"cabang" bson:"cabang"`
}
type Recap struct {
	Penjualan    []Penjualan   `json:"penjualan" bson:"penjualan"`
	Pengeluaran  []Pengeluaran `json:"pengeluaran" bson:"pengeluaran"`
	JumlahKotor  int           `json:"jumlahkotor" bson:"jumlahkotor"`
	JumlahBersih int           `json:"jumlahbersih" bson:"jumlahbersih"`
}
type User struct {
	ID       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
type Pengeluaran struct {
	ID              int    `json:"id" bson:"id"`
	Namapengeluaran string `json:"namapengeluaran" bson:"namapengeluaran"`
	Jumlah          int    `json:"jumlah" bson:"jumlah"`
	Tanggal         string `json:"tanggal" bson:"tanggal"`
	Cabang          string `json:"cabang" bson:"cabang"`
}

type RecapResponse struct {
	Penjualan         []Penjualan   `json:"penjualan" bson:"penjualan"`
	Pengeluaran       []Pengeluaran `json:"pengeluaran" bson:"pengeluaran"`
	JumlahKotor       string        `json:"jumlahkotor" bson:"jumlahkotor"`
	JumlahPengeluaran string        `json:"jumlahPengeluaran" bson:"jumlahPengeluaran"`
	JumlahBersih      string        `json:"jumlahbersih" bson:"jumlahbersih"`
}
