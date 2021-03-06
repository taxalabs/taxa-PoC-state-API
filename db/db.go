package db

import (
	"fmt"

	"github.com/3dw1nM0535/Byte/utils"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbHost, dbPort, dbUser, dbPassword, dbName, sslMode string
var logMode bool

func init() {
	godotenv.Load()
	dbHost = utils.MustGetEnv("DBHOST")
	dbPort = utils.MustGetEnv("DBPORT")
	dbUser = utils.MustGetEnv("DBUSER")
	dbPassword = utils.MustGetEnv("DBPASS")
	dbName = utils.MustGetEnv("DBNAME")
	sslMode = utils.MustGetEnv("SSLMODE_ENABLED")
	logMode = utils.MustGetEnvBool("LOGGING_STATUS")

}

// ORM : database structure
type ORM struct {
	DB *gorm.DB
}

// Factory : open a database connection
func Factory() (*ORM, error) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	// pass dialect and db URI to gorm.Open
	dbm, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Printf("Error connecting to database: " + err.Error())
	}
	dbm.LogMode(logMode) // should the db log every sql commands?(true or false)

	return &ORM{DB: dbm}, nil
}
