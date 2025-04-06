package video_domain

import (
	"youtubent/internal/db"

	"github.com/valyala/fasthttp"
)

type SqlcVideoRepository struct {
	db *db.Queries
}

func NewSqlcVideoRepository(db *db.Queries) *SqlcVideoRepository {
	return &SqlcVideoRepository{db: db}
}

func (r *SqlcVideoRepository) CreateVideo(ctx *fasthttp.RequestCtx, video *db.CreateVideoParams) error {
	_, err := r.db.CreateVideo(ctx, *video)
	return err
}