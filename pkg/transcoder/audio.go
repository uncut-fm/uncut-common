package transcoder

import (
	"errors"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
)

func (a API) GetAudioMetadataByURL(audioUrl string) (*model.AudioMetadata, error) {
	response, err := a.makeAudioMetadataRequest(audioUrl)
	if a.log.CheckError(err, a.GetImageMetadataByURL) != nil {
		return nil, err
	}

	return response, err
}

func (a API) makeAudioMetadataRequest(audioUrl string) (*model.AudioMetadata, error) {
	metadata := new(model.AudioMetadata)

	resp, err := a.restyClient.R().EnableTrace().
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
