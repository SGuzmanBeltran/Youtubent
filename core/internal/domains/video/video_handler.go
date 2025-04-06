package video_domain

import "github.com/gofiber/fiber/v2"

type CoreVideoHandler struct {
	videoService VideoService
}

func NewCoreVideoHandler(videoService VideoService) *CoreVideoHandler {
	return &CoreVideoHandler{videoService: videoService}
}

func (h *CoreVideoHandler) SetupRoutes(api fiber.Router) {
	api.Post("/video", h.CreateVideo)
}

func (h *CoreVideoHandler) CreateVideo(c *fiber.Ctx) error {
	var video *CreateVideo
	if err := c.BodyParser(&video); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return h.videoService.CreateVideo(c.Context(), video)
}
