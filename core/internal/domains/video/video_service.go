package video_domain

import (
	"github.com/valyala/fasthttp"
)

type CoreVideoService struct {
	videoRepository VideoRepository
}

func NewCoreVideoService(videoRepository VideoRepository) *CoreVideoService {
	return &CoreVideoService{videoRepository: videoRepository}
}

func (s *CoreVideoService) CreateVideo(ctx *fasthttp.RequestCtx, video *CreateVideo) error {
	dbVideo := MapCreateVideoToDbVideo(video)
	return s.videoRepository.CreateVideo(ctx, dbVideo)
}