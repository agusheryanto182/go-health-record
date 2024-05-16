package handler

import (
	"github.com/agusheryanto182/go-health-record/module/feature/image"
	"github.com/agusheryanto182/go-health-record/module/feature/image/svc"
	"github.com/agusheryanto182/go-health-record/utils/response"
	"github.com/gofiber/fiber/v2"
)

type ImageHandler struct {
	svc svc.ImageSvc
}

func NewImageHandler(svc svc.ImageSvc) image.ImageHandlerInterface {
	return &ImageHandler{
		svc: svc,
	}
}

func (c *ImageHandler) UploadImage(ctx *fiber.Ctx) error {
	fileHeader, err := ctx.FormFile("file")
	if fileHeader == nil {
		return response.NewBadRequestError("file should not be empty")
	}
	if err != nil {
		return response.NewInternalServerError("failed to retrieve file")
	}

	url, err := c.svc.UploadImage(fileHeader)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "image uploaded successfully",
		"data": fiber.Map{
			"imageUrl": url,
		},
	})

}
