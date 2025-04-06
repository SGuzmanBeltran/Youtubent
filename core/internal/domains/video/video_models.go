package video_domain

type CreateVideo struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Url         string  `json:"url"`
	Thumbnail   string  `json:"thumbnail"`
	Duration    int     `json:"duration"`
	Format      string  `json:"format"`
	Type        string  `json:"type"`
	Tags        []string `json:"tags"`
}
