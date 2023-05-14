package database

import (
	"fmt"
	"log"
	"sync"

	"teams/env"
	"teams/models"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	mu sync.Mutex
)

// migrations for database
func runMigrations(dsn string) {
	run, err := migrate.New(
		"file://./database/migrations",
		dsn,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := run.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println("Error running database migrations")
		log.Fatal(err)
	}

	sourceError, databaseError := run.Close()

	if sourceError != nil || databaseError != nil {
		if sourceError != nil {
			log.Println(sourceError)
		}

		if databaseError != nil {
			log.Println(databaseError)
		}

		log.Fatal("Error closing database connection for migrations")
	}
}

// opens connection to the database, returns an error if something wrong happens
func Open() error {
	host := env.GetOr("DB_HOST", "localhost")
	port := env.GetIntOr("DB_PORT", 5432)
	dbName := env.GetOr("DB_NAME", "teams")
	user := env.GetOr("DB_USER", "teams")
	password := env.GetOr("DB_PASSWORD", "teams")
	timezone := env.GetOr("TIMEZONE", "America/Denver")

	runMigrations(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbName))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", host, user, password, dbName, port, timezone)

	// open connections
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return err
	}

	return nil
}

func Close() {
	database, err := db.DB()

	if err != nil {
		log.Fatalf("Error getting database while closing: %v", err)
	}

	log.Println("Closing database connection")
	err = database.Close()

	if err != nil {
		log.Fatalf("Error closing database %v", err)
	}
}

func DB() *gorm.DB {

	return db
}

func Insert(user *models.User) {
	mu.Lock()
	db.Create(user)
	mu.Unlock()
}

func Get() []*models.User {
	var users []*models.User
	db.Find(&users)
	return users
}

// Delete a user from the database
func Delete(user *models.User) {
	mu.Lock()
	db.Delete(user)
	mu.Unlock()
}

// Add user authentication using common techniques
// such as hashing and salting passwords, JWT tokens, and middleware
// to protect routes that require authentication

// Add postgres database migrations
// using a migration tool such as golang-migrate
// to manage changes to the database schema over time.
// This ensures that the database schema is always up-to-date
// and that changes can be made safely without losing data.

// Example migration command:
// migrate -path database/migrations -database "postgresql://user:password@localhost:5432/dbname?sslmode=disable" up

// Note: The specific implementation of authentication and migrations
// will depend on the requirements of the application and the preferences of the developer.
// Add user authentication using common techniques
// such as hashing and salting passwords, JWT tokens, and middleware
// to protect routes that require authentication

// Example implementation using JWT tokens and bcrypt for password hashing:

// import (
//     "github.com/dgrijalva/jwt-go"
//     "golang.org/x/crypto/bcrypt"
// )

// // Define a struct to hold the user credentials
// type Credentials struct {
//     Username string `json:"username"`
//     Password string `json:"password"`
// }

// // Define a struct to hold the JWT token
// type Claims struct {
//     Username string `json:"username"`
//     jwt.StandardClaims
// }

// // Define a secret key for signing the JWT token
// var jwtKey = []byte("my_secret_key")

// // Authenticate a user by checking their credentials against the database
// func AuthenticateUser(username string, password string) bool {
//     var user models.User
//     db.Where("username = ?", username).First(&user)
//     if user.ID == 0 {
//         return false
//     }
//     err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
//     if err != nil {
//         return false
//     }
//     return true
// }

// // Generate a JWT token for a user
// func GenerateToken(username string) (string, error) {
//     // Set the expiration time for the token
//     expirationTime := time.Now().Add(24 * time.Hour)
//     // Create the JWT claims, which includes the username and expiration time
//     claims := &Claims{
//         Username: username,
//         StandardClaims: jwt.StandardClaims{
//             ExpiresAt: expirationTime.Unix(),
//         },
//     }
//     // Create the token using the claims and the secret key
//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//     tokenString, err := token.SignedString(jwtKey)
//     if err != nil {
//         return "", err
//     }
//     return tokenString, nil
// }

// // Middleware function to protect routes that require authentication
// func AuthMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Get the JWT token from the Authorization header
//         authHeader := r.Header.Get("Authorization")
//         if authHeader == "" {
//             http.Error(w, "Authorization header required", http.StatusUnauthorized)
//             return
//         }
//         tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
//         // Parse the JWT token
//         claims := &Claims{}
//         token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
//             return jwtKey, nil
//         })
//         if err != nil {
//             if err == jwt.ErrSignatureInvalid {
//                 http.Error(w, "Invalid token signature", http.StatusUnauthorized)
//                 return
//             }
//             http.Error(w, "Invalid token", http.StatusUnauthorized)
//             return
//         }
//         if !token.Valid {
//             http.Error(w, "Invalid token", http.StatusUnauthorized)
//             return
//         }
//         // Set the username in the request context
//         ctx := context.WithValue(r.Context(), "username", claims.Username)
//         // Call the next handler with the updated context
//         next.ServeHTTP(w, r.WithContext(ctx))
//     })
// }
