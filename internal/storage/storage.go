package storage

import (
	"fmt"
	"github.com/Grishun/kamen34/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Storage struct {
		logger *log.Logger
		db     *gorm.DB
		cfg    *Config
	}

	Config struct {
		Host     string
		User     string
		Password string
		DBName   string
		Port     uint
	}
)

func NewStorage(logger *log.Logger, config *Config) (_ *Storage, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Shanghai",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
	)

	logger.Printf("connected to data base with dsn %s \n", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Storage{
		logger: logger,
		db:     db,
		cfg:    config,
	}, nil

}

func (s *Storage) GetGORMDB() *gorm.DB {
	return s.db
}

func (s *Storage) MigrateUp() error {
	s.logger.Printf("try to migrate up psql")

	return s.db.AutoMigrate(&model.Product{})
}

func (s *Storage) CreateProduct(product *model.Product) error {
	s.logger.Println("try to create product")
	tx := s.db.Exec("INSERT INTO products(name, img_url, price) VALUES ($1,$2,$3);",
		product.Name,
		product.ImgURL,
		product.Price,
	)
	tx.Commit()
	return nil
}
