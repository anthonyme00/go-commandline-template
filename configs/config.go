package configs

import (
	"os"
	"reflect"
	"sync"

	"github.com/subosito/gotenv"
)

// GlobalConfig is the global configuration for the application.
// All fields are string only
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
		LoadFromEnv()
	})
}

func LoadFromEnv() {
	numField := reflect.TypeOf(Global).NumField()

	for i := 0; i < numField; i++ {
		field := reflect.TypeOf(Global).Field(i)
		env := field.Tag.Get("env")
		if env != "" {
			reflect.ValueOf(&Global).Elem().Field(i).SetString(os.Getenv(env))
		}
	}
}
