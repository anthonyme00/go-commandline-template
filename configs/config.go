package configs

import (
	"sync"

	"github.com/subosito/gotenv"
)

// DB_HOST = localhost
// DB_NAME = db
// DB_PORT = 3306
// DB_USER = USERNAME
// DB_PASS = PASSWORD

// MONGO_HOST = localhost
// MONGO_PORT = 27017
// MONGO_DB = DB
// MONGO_COLLECTION = COLLECTION

// ELASTIC_HOST = localhost
// ELASTIC_PORT = 39200
// ELASTIC_SCHEME = https
// ELASTIC_USERNAME = USERNAME
// ELASTIC_PASSWORD = PASSWORD
// ELASTIC_INDEX = INDEX
// ELASTIC_TYPE = TYPE
type GlobalConfig struct {
	PROJECT_PATH string

	DB_HOST string `env:"DB_HOST"`
	DB_NAME string `env:"DB_NAME"`
	DB_PORT string `env:"DB_PORT"`
	DB_USER string `env:"DB_USER"`
	DB_PASS string `env:"DB_PASS"`

	MONGO_HOST       string `env:"MONGO_HOST"`
	MONGO_PORT       string `env:"MONGO_PORT"`
	MONGO_DB         string `env:"MONGO_DB"`
	MONGO_COLLECTION string `env:"MONGO_COLLECTION"`

	ELASTIC_HOST     string `env:"ELASTIC_HOST"`
	ELASTIC_PORT     string `env:"ELASTIC_PORT"`
	ELASTIC_SCHEME   string `env:"ELASTIC_SCHEME"`
	ELASTIC_USERNAME string `env:"ELASTIC_USERNAME"`
	ELASTIC_PASSWORD string `env:"ELASTIC_PASSWORD"`
	ELASTIC_INDEX    string `env:"ELASTIC_INDEX"`
	ELASTIC_TYPE     string `env:"ELASTIC_TYPE"`
}

var Global GlobalConfig
var once sync.Once

func Init() {
	once.Do(func() {
		gotenv.Load("./.env")
		Global = GlobalConfig{
			PROJECT_PATH: "./",
		}
	})
}
