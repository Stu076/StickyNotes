package database

import (
	supa "github.com/nedpals/supabase-go"
	"github.com/spf13/viper"
)

type DB struct {
	Client *supa.Client
}

func New() *DB {
	return &DB{
		Client: initDB(),
	}
}

func initDB() *supa.Client {
	baseUrl := viper.GetString("database.supabaseUrl")
	baseKey := viper.GetString("database.supabaseKey")

	client := supa.CreateClient(baseUrl, baseKey)

	return client
}
