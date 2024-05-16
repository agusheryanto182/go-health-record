package routes

import (
	"github.com/agusheryanto182/go-health-record/module/feature/image"
	"github.com/agusheryanto182/go-health-record/module/feature/medical"
	"github.com/agusheryanto182/go-health-record/module/feature/user"
	"github.com/agusheryanto182/go-health-record/module/middleware"
	"github.com/agusheryanto182/go-health-record/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, handler user.UserHandlerInterface, jwtSvc jwt.JWTInterface, userSvc user.UserSvcInterface) {
	user := app.Group("/v1/user")
	user.Post("/it/register", handler.RegisterIt)
	user.Post("/it/login", handler.LoginIt)
	user.Post("/nurse/login", handler.LoginNurse)
	user.Post("/nurse/register", middleware.Protected(jwtSvc, userSvc), handler.RegisterNurse)
	user.Get("", middleware.Protected(jwtSvc, userSvc), handler.GetUserByFilters)
	user.Put("/nurse/:userId", middleware.Protected(jwtSvc, userSvc), handler.UpdateUserNurse)
	user.Delete("/nurse/:userId", middleware.Protected(jwtSvc, userSvc), handler.DeleteUserNurse)
	user.Post("/nurse/:userId/access", middleware.Protected(jwtSvc, userSvc), handler.SetPasswordNurse)
}

func MedicalRoute(app *fiber.App, handler medical.MedicalHandlerInterface, jwtSvc jwt.JWTInterface, userSvc user.UserSvcInterface) {
	medical := app.Group("/v1/medical")
	medical.Post("/record", middleware.Protected(jwtSvc, userSvc), handler.CreateRecord)
	medical.Get("/record", middleware.Protected(jwtSvc, userSvc), handler.GetRecordByFilters)
	medical.Get("/patient", middleware.Protected(jwtSvc, userSvc), handler.GetPatientByFilters)
	medical.Post("/patient", middleware.Protected(jwtSvc, userSvc), handler.RegisterPatient)
}

func ImageRoute(app *fiber.App, handler image.ImageHandlerInterface, jwtSvc jwt.JWTInterface, userSvc user.UserSvcInterface) {
	image := app.Group("/v1/image")
	image.Post("", middleware.Protected(jwtSvc, userSvc), handler.UploadImage)
}
