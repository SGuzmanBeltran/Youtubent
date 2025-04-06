package video_domain

import (
	"youtubent/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// Repository
type VideoRepository interface {
	CreateVideo(ctx *fasthttp.RequestCtx, video *db.CreateVideoParams) error
}

// Service
type VideoService interface {
	CreateVideo(ctx *fasthttp.RequestCtx, video *CreateVideo) error
}

// Handler
type VideoHandler interface {
	SetupRoutes(app *fiber.App)
	CreateVideo(c *fiber.Ctx) error
}
