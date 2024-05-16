package main

import (
	"log"

	"github.com/agusheryanto182/go-health-record/config"
	userHandler "github.com/agusheryanto182/go-health-record/module/feature/user/handler"
	userRepo "github.com/agusheryanto182/go-health-record/module/feature/user/repo"
	userSvc "github.com/agusheryanto182/go-health-record/module/feature/user/svc"
	"github.com/agusheryanto182/go-health-record/module/middleware"
	"github.com/agusheryanto182/go-health-record/routes"
	"github.com/agusheryanto182/go-health-record/utils/database"
	"github.com/agusheryanto182/go-health-record/utils/hash"
	"github.com/agusheryanto182/go-health-record/utils/jwt"
	"github.com/agusheryanto182/go-health-record/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
		AppName:      "Project Sprint Week 3 - Go Health Record API",
	})

	bootConfig := config.NewConfig()
	hash := hash.NewHash(bootConfig)
	jwt := jwt.NewJWTService(bootConfig)
	valid := validator.New()

	db, err := database.InitDatabase(bootConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// register validation
	valid.RegisterValidation("ValidateNipIt", validation.ValidateNipIt)
	valid.RegisterValidation("ValidateNipNurse", validation.ValidateNipNurse)
	valid.RegisterValidation("ValidateImage", validation.ValidateImage)
	valid.RegisterValidation("ValidatePhoneNumber", validation.ValidatePhoneNumberFormat)
	valid.RegisterValidation("ValidateURL", validation.ValidateURL)

	app.Use(recover.New())
	app.Use(middleware.Logger())

	// repo
	userRepo := userRepo.NewUserRepository(db)

	// svc
	userSvc := userSvc.NewUserService(userRepo, valid, hash, jwt)

	// handler
	userHandler := userHandler.NewUserHandler(userSvc)

	// routes
	routes.UserRoute(app, userHandler, jwt, userSvc)

	log.Fatal(app.Listen(":8080"))
}
