package repository

import (
	"context"
	"fmt"

	"github.com/gocroot/kampus/model"
	"github.com/rofinafiin/lapak-UMK/controller/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertPenjualan(db *mongo.Database, id string, namaproduk string, jumlahpenjualan string, tanggaldatamasuk string) (InsertedID interface{}, err error) {
	var m model.Penjualan
	m.ID = id
	m.NamaProduk = namaproduk
	m.JumlahPenjualan = jumlahpenjualan
	m.TanggalDataMasuk = tanggaldatamasuk
	return InsertOneDoc(db, "laporankeungan", m), err
}

func GetPenjualanByNamaProduk(namaproduk string, db *mongo.Database) (data model.Penjualan, err error) {
	user := db.Collection("laporankeungan")
	filter := bson.M{"NamaProduk": namaproduk}
	err = user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetDataCompFromStatus: %v\n", err)
	}
	return data, err
}
func InsertPengeluaran(db *mongo.Database, id string, nama string, jumlah string, tanggal string) (InsertedID interface{}, err error) {
	var m model.Penjualan
	m.ID = id
	m.Nama = nama
	m.Jumlah = jumlah
	m.Tanggal = tanggal
	return InsertOneDoc(db, "laporankeungan", m), err
}
func GetPengeluaranByNama(nama string, db *mongo.Database) (data model.Pengeluaran, err error) {
	user := db.Collection("laporankeungan")
	filter := bson.M{"Nama": nama}
	err = user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetDataCompFromStatus: %v\n", err)
	}
	return data, err
}
