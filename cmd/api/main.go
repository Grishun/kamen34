package main

import (
	"github.com/Grishun/kamen34/internal/model"
	"github.com/Grishun/kamen34/internal/storage"
	log2 "log"
)

func main() {
	log := log2.Default()
	cfg := &storage.Config{
		Host:     "127.0.0.1",
		User:     "grishu",
		Password: "postgres",
		DBName:   "kamen34",
		Port:     5432,
	}
	strg, err := storage.NewStorage(log, cfg)

	if err != nil {
		panic(err)
	}

	err = strg.MigrateUp()

	if err != nil {
		panic(err)
	}

	err = strg.CreateProduct(&model.Product{
		Name:   "grishu",
		ImgURL: "https://www.4-7.ru/netcat_files/30/34/Glyby_i_krupnye_valuny_Glyba_Shungit_10.jpg",
		Price:  8841,
	})

	if err != nil {
		panic(err)
	}

}
