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
// @Param namaproduk path string true "Masukan namaproduk"
// @Success 200 {object} model.Pengeluaran{}
// @Router /lapumk/pengeluaran/{namaproduk} [get]
func (db *UMKHandler) GetDataPengeluaran(c *fiber.Ctx) (err error) {
	namaproduk := c.Params("namaproduk")
	getdata, err := repository.GetPenjualanByNamaProduk(namaproduk, config.DBMongo("lapak-UMK"))
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
// @Param namapenjualan path string true "Masukan namapenjualan"
// @Success 200 {object} model.Penjualan{}
// @Router /lapumk/penjualan/{namapenjualan} [get]
func (db *UMKHandler) GetDataPenjualan(c *fiber.Ctx) (err error) {
	namapenjualan := c.Params("namapenjualan")
	getdata, err := repository.GetPenjualanByNamaProduk(namapenjualan, config.DBMongo("lapak-UMK"))
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
