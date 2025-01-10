package transcoder

import (
	"context"
	"errors"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
)

func (a API) GetAudioMetadataByURL(ctx context.Context, audioUrl string) (*model.AudioMetadata, error) {
	response, err := a.makeAudioMetadataRequest(ctx, audioUrl)
	if a.log.CheckError(err, a.GetImageMetadataByURL) != nil {
		return nil, err
	}

	return response, err
}

func (a API) makeAudioMetadataRequest(ctx context.Context, audioUrl string) (*model.AudioMetadata, error) {
	metadata := new(model.AudioMetadata)

	resp, err := a.restyClient.R().SetContext(ctx).
		SetHeader("admin-token", a.adminToken).
		SetResult(metadata).
		SetQueryParam("url", audioUrl).
		Get(fmt.Sprintf(audioMetadataEndpoint, a.url))

	if a.log.CheckError(err, a.makeAudioMetadataRequest) != nil {
		return nil, err
	}

	if resp.StatusCode() >= 300 || resp.StatusCode() < 200 {
		return nil, errors.New("failed getting video metadata")
	}

	return metadata, a.log.CheckError(err, a.makeAudioMetadataRequest)
}
