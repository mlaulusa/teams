package middleware

import (
	"teams/session"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func AddStandard(app *fiber.App) {

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(csrf.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(cache.New())
	app.Use(limiter.New(limiter.Config{
		Max: 20,
	}))
	app.Use(cors.New())
	// app.Use(secure.New())addProtected(app)
	// app.Use(ipware.New())
	// app.Use(rewrite.New())
	// app.Use(forwarded.New())
	app.Use(pprof.New())

	app.Use(session.New(session.Config{}))
}
