package component

import (
	"fmt"
	"log"
	"ridhoandhika/backend-api/domain"
	"ridhoandhika/backend-api/internal/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabaseConnection(cnf *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable",
		cnf.Database.Host,
		cnf.Database.Port,
		cnf.Database.User,
		cnf.Database.Password,
		cnf.Database.Name)

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Mengecek apakah koneksi dapat di-"ping"
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("error open connection %v", err.Error())
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("error open connection %v", err.Error())
	}
	db.Exec("CREATE SCHEMA IF NOT EXISTS custom_schema")
	// Melakukan migrasi ke database (membuat tabel user jika belum ada)
	err = db.Debug().AutoMigrate(&domain.User{}, &domain.PersonalInformation{}, &domain.WorkExperience{}, &domain.Education{})
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi: %v", err)
	}

	return db
}
