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
	ID              int    `json:"ID" bson:"ID"`
	Namapengeluaran string `json:"namapengeluaran" bson:"namapengeluaran"`
	Jumlah          int    `json:"jumlah" bson:"jumlah"`
	Tanggal         string `json:"tanggal" bson:"tanggal"`
	Cabang          string `json:"cabang" bson:"cabang"`
}
