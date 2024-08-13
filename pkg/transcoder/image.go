package transcoder

import (
	"errors"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
)

func (a API) GetImageMetadataByURL(imageURL string) (*model.ImageMetadata, error) {
	response, err := a.makeImageMetadataRequest(imageURL)
	if a.log.CheckError(err, a.GetImageMetadataByURL) != nil {
		return nil, err
	}

	return response, err
}

func (a API) makeImageMetadataRequest(imageURL string) (*model.ImageMetadata, error) {
	metadata := new(model.ImageMetadata)

	resp, err := a.restyClient.R().EnableTrace().
		SetHeader("admin-token", a.adminToken).
		SetResult(metadata).
		SetQueryParam("url", imageURL).
		Get(fmt.Sprintf(imageMetadata, a.url))

	if a.log.CheckError(err, a.makeImageMetadataRequest) != nil {
		return nil, err
	}

	if resp.StatusCode() >= 300 || resp.StatusCode() < 200 {
		return nil, errors.New("failed getting image metadata")
	}

	return metadata, a.log.CheckError(err, a.makeImageMetadataRequest)
}
