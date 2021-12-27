package pubsub

type CutAudioRequest struct {
	EpisodeAudioURL string `json:"episode_audio_url,required"`
	EpisodeID       int    `json:"episode_id,required"`
	MomentID        int    `json:"moment_id,required"`
	From            int    `json:"from,required"`
	To              int    `json:"to,required"`
}
