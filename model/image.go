package model

type ImageMetadata struct {
	Width            int
	Height           int
	AspectRatio      float64
	Size             int
	MimeType         string
	GifFirstFrameUrl *string
	DominantColorHSL *string
}

type VideoMetadata struct {
	Width         int     `json:"width"`
	Height        int     `json:"height"`
	AspectRatio   float64 `json:"aspect_ratio"`
	Size          int     `json:"size"`
	DurationInSec int     `json:"duration_in_sec"`
	MimeType      string  `json:"mime_type"`
	FirstFrameUrl *string `json:"first_frame_url"`
}
