package database

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/gabefgonc/PetShopMGMT/internal/app/infra/config"
)

var (
	db   *pgxpool.Pool
	once sync.Once
)

func NewPostgresDatabase(config *config.Config) *pgxpool.Pool {
	once.Do(func() {
		connUrl := fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
		)

		poolConfig, err := pgxpool.ParseConfig(connUrl)
		if err != nil {
			log.Fatalln("Unable to parse connection URL: ", err)
		}

		db, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err != nil {
			log.Fatalln("Unable to create PGX connection pool: ", err)
		}

		if err := db.Ping(context.Background()); err != nil {
			log.Fatalln("Failed to connect to database: ", err)
		}
	})

	return db
}
