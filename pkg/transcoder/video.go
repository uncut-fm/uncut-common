package transcoder

import (
	"context"
	"errors"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
)

func (a API) GetFirstFrameFromVideo(ctx context.Context, videoURL string) (string, error) {
	response, err := a.makeVideoFirstFrameRequest(ctx, videoURL)
	if a.log.CheckError(err, a.GetFirstFrameFromGif) != nil {
		return "", err
	}

	return response, err
}

func (a API) makeVideoFirstFrameRequest(ctx context.Context, videoUrl string) (string, error) {
	var result struct {
		Url string `json:"url"`
	}
	resp, err := a.restyClient.R().SetContext(ctx).
		SetHeader("admin-token", a.adminToken).
		SetResult(&result).
		SetQueryParam("url", videoUrl).
		Get(fmt.Sprintf(videoFirstFrame, a.url))

	if a.log.CheckError(err, a.makeVideoFirstFrameRequest) != nil {
		return "", err
	}

	if resp.StatusCode() >= 300 || resp.StatusCode() < 200 {
		return "", errors.New("failed getting video first frame")
	}

	return result.Url, a.log.CheckError(err, a.makeVideoFirstFrameRequest)
}

func (a API) GetVideoMetadataByURL(ctx context.Context, videoUrl string) (*model.VideoMetadata, error) {
	response, err := a.makeVideoMetadataRequest(ctx, videoUrl)
	if a.log.CheckError(err, a.GetImageMetadataByURL) != nil {
		return nil, err
	}

	return response, err
}

func (a API) makeVideoMetadataRequest(ctx context.Context, imageURL string) (*model.VideoMetadata, error) {
	metadata := new(model.VideoMetadata)

	resp, err := a.restyClient.R().SetContext(ctx).
		SetHeader("admin-token", a.adminToken).
		SetResult(metadata).
		SetQueryParam("url", imageURL).
		Get(fmt.Sprintf(videoMetadata, a.url))

	if a.log.CheckError(err, a.makeVideoMetadataRequest) != nil {
		return nil, err
	}

	if resp.StatusCode() >= 300 || resp.StatusCode() < 200 {
		return nil, errors.New("failed getting video metadata")
	}

	return metadata, a.log.CheckError(err, a.makeVideoMetadataRequest)
}
