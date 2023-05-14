package handlers

import (
	"log"
	"teams/database"
	"teams/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/x/crypto/bcrypt"
	"golang.org/x/crypto/bcrypt"
)

func addProtected(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Get("/hello", helloWorld)

	// Bind handlers
	v1.Get("/users", userList)
	v1.Post("/users", userCreate)

	v1.Post("/test", test)
}

func Add(app *fiber.App) {
	addProtected(app)
	// send client files
	app.Static("/", "./static/public")

	app.Use(notFound)
}

func helloWorld(c *fiber.Ctx) error {

	return c.SendString("Hello World")
}

func test(c *fiber.Ctx) error {
	return c.SendString("Computer guy")
}

// UserList returns a list of users
func userList(c *fiber.Ctx) error {
	users := database.Get()

	return c.JSON(fiber.Map{
		"success": true,
		"users":   users,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// UserCreate registers a user
func userCreate(c *fiber.Ctx) error {

	PasswordHash, err := HashPassword(c.FormValue("password"))

	if err != nil {
		log.Fatalf(err.Error())
	}

	user := &models.User{
		// Note: when writing to external database,
		// we can simply use - Name: c.FormValue("user")
		FirstName:    utils.CopyString(c.FormValue("first_name")),
		LastName:     utils.CopyString(c.FormValue("last_name")),
		Email:        utils.CopyString(c.FormValue("email")),
		PasswordHash: PasswordHash,
	}

	database.Insert(user)

	return c.JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
}

// NotFound returns custom 404 page
func notFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
