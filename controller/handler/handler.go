package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rofinafiin/lapak-UMK/config"
	"github.com/rofinafiin/lapak-UMK/controller/helper/json"
	"github.com/rofinafiin/lapak-UMK/controller/model"
	"github.com/rofinafiin/lapak-UMK/controller/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type UMKHandler struct {
	Mongo *mongo.Database
}

// GetDataPengeluaran godoc
// @Summary Mengambil data Pengeluaran (single).
// @Description get data pengeluaran.
// @Tags Lapak-UMKM
// @Accept application/json
// @Produce json
// @Param namapengeluaran path string true "Masukan namapengeluaran"
// @Success 200 {object} model.Pengeluaran{}
// @Router /lapumk/pengeluaran/{namapengeluaran} [get]
func (db *UMKHandler) GetDataPengeluaran(c *fiber.Ctx) (err error) {
	namapenjualan := c.Params("namapengeluaran")
	getdata, err := repository.GetPengeluaranByNama(namapenjualan, config.DBMongo("lapak-UMK"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data Tidak ada")
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Pengeluaran berhasil diambil",
		Data:    getdata,
	}.WriteToBody(c)
}

// GetDataPenjualan godoc
// @Summary Mengambil data Penjualan (single).
// @Description get data Penjualan.
// @Tags Lapak-UMKM
// @Accept application/json
// @Produce json
// @Param NamaProduk path string true "Masukan namaproduk"
// @Success 200 {object} model.Penjualan{}
// @Router /lapumk/penjualan/{NamaProduk} [get]
func (db *UMKHandler) GetDataPenjualan(c *fiber.Ctx) (err error) {
	namaproduk := c.Params("NamaProduk")
	getdata, err := repository.GetPenjualanByNamaProduk(namaproduk, config.DBMongo("lapak-UMK"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data tidak Ditemukan")
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Penjualan berhasil diambil",
		Data:    getdata,
	}.WriteToBody(c)
}

// InsertDataPenjualan godoc
// @Summary insert data penjualan.
// @Description get data penjualan.
// @Tags Lapak-UMKM
// @Accept application/json
// @Param request body model.Penjualan true "Payload Body [RAW]"
// @Produce json
// @Success 200 {object} model.Penjualan
// @Router /lapumk/inspenjualan [post]
func (db *UMKHandler) InsertDataPenjualan(c *fiber.Ctx) (err error) {
	database := config.DBMongo("lapak-UMK")
	var penjualan model.Penjualan
	if err := c.BodyParser(&penjualan); err != nil {
		return err
	}
	Inserted, err := repository.InsertPenjualan(database,
		penjualan.ID,
		penjualan.NamaProduk,
		penjualan.JumlahPenjualan,
		penjualan.TanggalDataMasuk,
		penjualan.Cabang,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Berhasil Disimpan",
		Data:    Inserted,
	}.WriteToBody(c)
}

// InsPengeluaran godoc
// @Summary insert data Pengeluaran.
// @Description get data Pengeluaran.
// @Tags Lapak-UMKM
// @Accept application/json
// @Param request body model.Pengeluaran true "Payload Body [RAW]"
// @Produce json
// @Success 200 {object} model.Pengeluaran
// @Router /lapumk/inspengeluaran [post]
func (db *UMKHandler) InsPengeluaran(c *fiber.Ctx) (err error) {
	database := config.DBMongo("lapak-UMK")
	var pengeluaran model.Pengeluaran
	if err := c.BodyParser(&pengeluaran); err != nil {
		return err
	}
	Inserted, err := repository.InsertPengeluaran(database,
		pengeluaran.ID,
		pengeluaran.Namapengeluaran,
		pengeluaran.Jumlah,
		pengeluaran.Tanggal,
		pengeluaran.Cabang,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Berhasil Disimpan",
		Data:    Inserted,
	}.WriteToBody(c)
}

func (db *UMKHandler) KalkulasiLaporan(c *fiber.Ctx) (err error) {
	return err
}

func (db *UMKHandler) GetAllPengeluaran(c *fiber.Ctx) (err error) {

	return err
}
