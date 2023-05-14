package database

import (
	"log"
	"warehouse/database/credentials"
	"warehouse/database/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func Connect() {
	url := "host=" + credentials.DB_HOST + " user=" + credentials.DB_USERNAME + " password=" + credentials.DB_PASSWORD + " dbname=" + credentials.DB_DATABASE + " port=" + credentials.DB_PORT

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Fatal("Something went wrong during the connection with the database.", err)
	}

	log.Println("Successfully connected to the database.")

	// Configure the database driver
	driver, _ := db.DB()

	driver.SetMaxIdleConns(10)
	driver.SetMaxOpenConns(100)

	// Migrate the schema
	db.AutoMigrate(&entities.Role{})
	db.AutoMigrate(&entities.TFA{})
	db.AutoMigrate(&entities.Key{})
	db.AutoMigrate(&entities.App{})
	db.AutoMigrate(&entities.Account{})
	db.AutoMigrate(&entities.Action{})
	db.AutoMigrate(&entities.Changelog{})
	db.AutoMigrate(&entities.Version{})
	db.AutoMigrate(&entities.Branch{})

	// Create default roles
	if db.Where("name = ?", "Administrator").First(&entities.Role{}).RowsAffected == 0 {
		db.Create(&entities.Role{
			Name:            "Default",
			PermissionLevel: 0,
		})

		db.Create(&entities.Role{
			Name:            "Administrator",
			PermissionLevel: 100,
		})
	}

	// Assign the database to the global variable
	DBConn = db
}

func DefaultRole() (entities.Role, error) {

	var role entities.Role
	if err := DBConn.Where("name = ?", "Default").First(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}
