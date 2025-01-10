package transcoder

import (
	"context"
	"errors"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
)

func (a API) GetImageMetadataByURL(ctx context.Context, imageURL string) (*model.ImageMetadata, error) {
	response, err := a.makeImageMetadataRequest(ctx, imageURL)
	if a.log.CheckError(err, a.GetImageMetadataByURL) != nil {
		return nil, err
	}

	return response, err
}

func (a API) makeImageMetadataRequest(ctx context.Context, imageURL string) (*model.ImageMetadata, error) {
	metadata := new(model.ImageMetadata)

	resp, err := a.restyClient.R().SetContext(ctx).
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
