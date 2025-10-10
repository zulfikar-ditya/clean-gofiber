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

	app := fiber.New(fiber.Config{
		AppName: appConfig.APP_NAME,
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

	log.Println("Redis connected successfully")

	routes.SetupRoutes(app)

	app.Listen(":" + appConfig.APP_PORT)
}