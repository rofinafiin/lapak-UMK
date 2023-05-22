package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rofinafiin/lapak-UMK/config"
	"github.com/rofinafiin/lapak-UMK/controller/handler"
)

func SetuplapRoutes(router fiber.Router) {

	Mongo := config.DBMongo("lapak-UMK")

	UMK := handler.UMKHandler{Mongo}

	router.Get("/penjualan/:namaproduk", UMK.GetDataPenjualan)
	router.Get("/pengeluaran/:namapengeluaran", UMK.GetDataPengeluaran)
	router.Post("/inspenjualan", UMK.InsertDataPenjualan)

}
