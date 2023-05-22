package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rofinafiin/lapak-UMK/config"
	"github.com/rofinafiin/lapak-UMK/controller/helper/json"
	"github.com/rofinafiin/lapak-UMK/controller/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type UMKHandler struct {
	Mongo *mongo.Database
}

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
