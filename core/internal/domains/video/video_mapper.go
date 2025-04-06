package video_domain

import (
	"youtubent/internal/db"

	"github.com/jackc/pgx/v5/pgtype"
)

func MapCreateVideoToDbVideo(video *CreateVideo) *db.CreateVideoParams {
	return &db.CreateVideoParams{
		Title:           video.Title,
		Description:     pgtype.Text{String: video.Description, Valid: true},
		ThumbnailPath:   video.Thumbnail,
		Path:            video.Url,
		DurationSeconds: int32(video.Duration),
		Format:          video.Format,
		Type:            video.Type,
		Tags:            video.Tags,
	}
}
