package db

import (
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB_URI = os.Getenv("DB_URI")

func InitMGM() error {
	if DB_URI == "" {
		log.Println("No DB_URI provided, using localhost instead")
		DB_URI = "mongodb://localhost:27017"
	}
	log.Println("⏳ Initializing MGM ... 🗺")
	err := mgm.SetDefaultConfig(nil, "prosperity-game", options.Client().ApplyURI(DB_URI))
	if err != nil {
		log.Println("❌ Failed to connect to DB")
		panic(err)
	}
	log.Println("🔰 Looks like MGM is initialized")
	return err
}
