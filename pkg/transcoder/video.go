package transcoder

import (
	"errors"
	"fmt"
)

func (a API) GetFirstFrameFromVideo(videoURL string) (string, error) {
	response, err := a.makeVideoFirstFrameRequest(videoURL)
	if a.log.CheckError(err, a.GetFirstFrameFromGif) != nil {
		return "", err
	}

	return response, err
}

func (a API) makeVideoFirstFrameRequest(videoUrl string) (string, error) {
	var result struct {
		Url string `json:"url"`
	}
	resp, err := a.restyClient.R().EnableTrace().
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
