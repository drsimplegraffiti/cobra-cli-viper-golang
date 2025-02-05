package utils

import (
	"database/sql"
	"fmt"
	"gonote/internal/store"
	"log"
	"sync"

	"github.com/spf13/viper"
	_ "github.com/lib/pq"
)

// Singleton pattern to reuse the DB connection
var (
	dbInstance *store.Queries
	once       sync.Once
)

// GetDB initializes and returns a singleton instance of SQLC queries
func GetDB() *store.Queries {
	once.Do(func() {
		// Read database config from Viper
		dbHost := viper.GetString("database.host")
		dbPort := viper.GetInt("database.port")
		dbUser := viper.GetString("database.user")
		dbPassword := viper.GetString("database.password")
		dbName := viper.GetString("database.dbname")
		sslMode := viper.GetString("database.sslmode")

		// Construct the PostgreSQL connection string
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
			dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

		// Open the database connection
		conn, err := sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Initialize SQLC Queries instance
		dbInstance = store.New(conn)
	})
	return dbInstance
}
