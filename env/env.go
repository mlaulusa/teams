package env

// Import the necessary packages
import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	REDIS   string = "STORE"
	SESSION string = "SESSION"
	ACCOUNT string = "ACCOUNT"
)

func init() {
	fmt.Println("Loading environment variables")

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func Get(key string) (string, bool) {
	variable, ok := os.LookupEnv(key)

	return variable, ok
}

func GetOr(key, defaultValue string) string {
	variable, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	return variable
}

func GetIntOr(key string, defaultValue int) int {
	variable, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	num, err := strconv.Atoi(variable)

	if err != nil {
		log.Printf("Cannot turn environment variable found with key %s into a number", key)
		log.Println(err)
		return defaultValue
	}

	return num
}
