package db

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Trivy []struct {
	Target          string
	Vulnerabilities []struct {
		VulnerabilityID  string
		PkgName          string
		InstalledVersion string
		FixedVersion     string
		Title            string
		Description      string
		Severity         string
		References       interface{}
	}
}

type Dockle struct {
	Summary struct {
		Fatal int
		Warn  int
		Info  int
		Pass  int
	}
	Details []struct {
		Code   string
		Title  string
		Level  string
		Alerts []string
	}
}

type Coninfo struct {
	Id          int `gorm:"primary_key";"AUTO_INCREMENT"`
	CreatedTime time.Time
	UpdatedTime time.Time
	ImageID     string
	RepoTag     string
	HostName    string
	ToolVersion string
	TrivyData   *Trivy
	DockleData  *Dockle
}

var db *gorm.DB
var err error

func Connect() {

	// Need to take values from ENV using os.GetEnv() if not passed take default values
	// connStr := os.ExpandEnv("host=${DB_HOST} user=${DB_USER} dbname=${DB_NAME} sslmode=disable password=${DB_PASSWORD}"))
	connStr := "host=localhost port=5432 user=coolman password=XXXXXXXXXXXXXXXXXX dbname=consecmon sslmode=disable"

	db, err := gorm.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	log.Printf("Connected to database successfully!\n")

	db.CreateTable(&Coninfo{})

	// db.Debug().AutoMigrate(&Coninfo{})
	db.AutoMigrate(&Coninfo{})

}
