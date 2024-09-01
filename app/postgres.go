package main

import (
	"go/djan/app/ent"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	clientInstance *ent.Client
	once           sync.Once
)

// GetClient returns the singleton instance of the ent.Client
func GetClient() *ent.Client {
	once.Do(func() {
		var err error
		user := viper.GetString("POSTGRES_USER")
		db_name := viper.GetString("POSTGRES_DBNAME")
		password := viper.GetString("POSTGRES_PASSWORD")
		db_string := "host=localhost port=5432 user=" + user + " dbname=" + db_name + " password=" + password
		clientInstance, err = ent.Open("postgres", db_string)
		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}
	})
	return clientInstance
}
