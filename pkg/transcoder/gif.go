package transcoder

import (
	"errors"
	"fmt"
)

func (a API) GetFirstFrameFromGif(gifURL string) (string, error) {
	response, err := a.makeGifFirstFrameRequest(gifURL)
	if a.log.CheckError(err, a.GetFirstFrameFromGif) != nil {
		return "", err
	}

	return response, err
}

func (a API) makeGifFirstFrameRequest(gifURL string) (string, error) {
	var result struct {
		Url string `json:"url"`
	}
	resp, err := a.restyClient.R().EnableTrace().
		SetHeader("admin-token", a.adminToken).
		SetResult(&result).
		SetQueryParam("url", gifURL).
		Get(fmt.Sprintf(gifFirstFrame, a.url))

	if a.log.CheckError(err, a.makeGifFirstFrameRequest) != nil {
		return "", err
	}

	if resp.StatusCode() >= 300 || resp.StatusCode() < 200 {
		return "", errors.New("failed getting gif first frame")
	}

	return result.Url, a.log.CheckError(err, a.makeGifFirstFrameRequest)
}
