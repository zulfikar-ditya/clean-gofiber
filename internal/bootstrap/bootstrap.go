package bootstrap

import (
	"aolus-software/clean-gofiber/adapters/http/routes"
	"aolus-software/clean-gofiber/internal/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AppBootstrap() {
	viper := config.NewViper()
	appConfig := config.AppConfig(viper)

	errorHandler := func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		// ignore other error (404, 403, 405, etc) only log 500 errors
		if code == fiber.StatusInternalServerError {
			log.Printf("Fiber error: %v", err)
		}

		return c.Status(code).JSON(fiber.Map{
			"error":   err.Error(),
			"status":  code,
		})
	}

	app := fiber.New(fiber.Config{
		AppName:      appConfig.APP_NAME,
		ErrorHandler: errorHandler,
	})

	databaseConfig := config.NewDatabaseConfig(viper)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  databaseConfig.GetDSN(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database instance: " + err.Error())
	}
	
	if err := sqlDB.Ping(); err != nil {
		panic("Failed to ping database: " + err.Error())
	}

	redisConfig := config.NewRedisConfig(viper)
	redisClient := redisConfig.NewRedisClient()
	defer redisClient.Close()

	routes.SetupRoutes(app)

	app.Listen(":" + appConfig.APP_PORT)
}