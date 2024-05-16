package image

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type ImageSvcInterface interface {
	UploadImage(*multipart.FileHeader) (string, error)
}

type ImageHandlerInterface interface {
	UploadImage(ctx *fiber.Ctx) error
}
