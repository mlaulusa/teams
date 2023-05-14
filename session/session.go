package session

import (
	"log"
	"runtime"
	"time"

	"teams/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v2"
)

type Config struct {
}

type Store = *session.Store

var (
	storage *redis.Storage
	store   *session.Store
)

func openRedis() *redis.Storage {

	host := env.GetOr("REDIS_HOST", "localhost")
	port := env.GetIntOr("REDIS_PORT", 6379)
	db := env.GetIntOr("REDIS_DB", 0)
	username := env.GetOr("REDIS_USERNAME", "")
	password := env.GetOr("REDIS_PASSWORD", "")

	redisStorage := redis.New(redis.Config{
		Host:      host,
		Port:      port,
		Username:  username,
		Password:  password,
		Database:  db,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})

	return redisStorage
}

func Start() {
	storage = openRedis()

	store = session.New(session.Config{
		Expiration: 4 * time.Hour,
		Storage:    storage,
	})
}

func Close() {
	log.Println("Closing redis connection")

	err := storage.Close()

	if err != nil {
		log.Fatalf("Error closing redis connection: %v", err)
	}
}

func GetSession() *session.Store {
	return store
}

func New(config Config) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Locals(env.SESSION, store)

		return c.Next()
	}
}
