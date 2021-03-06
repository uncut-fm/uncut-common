package storage

import (
	"context"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
	"time"
)

func (s Client) storeNftFile(c context.Context, nftID *int, file []byte, extension string, nftFilePath func(nftID *int, extension *string) string) (string, error) {
	fileName := nftFilePath(nftID, &extension)

	err := s.uploadFile(c, fileName, file)
	if err != nil {
		return "", err
	}

	err = s.MakeFilePublic(c, fileName)
	if err != nil {
		return "", err
	}

	return GetPublicFilePath(s.bucket, fileName), nil
}

func (s Client) getNftVideoFilepath(nftId *int, extension *string) string {
	now := time.Now()

	var fileName string
	if model.IsIntNil(nftId) {
		fileName = fmt.Sprintf(nftVideoFileFormat, now.Unix())
	} else {
		fileName = fmt.Sprintf(nftVideoIDFileFormat, *nftId, now.Unix())
	}

	if !model.IsStringNil(extension) {
		fileName = fmt.Sprintf("%s.%s", fileName, *extension)
	}

	return GetNftLocationPath(s.environment, fileName)
}

func (s Client) getNftImageFilepath(nftId *int, extension *string) string {
	now := time.Now()

	var fileName string
	if model.IsIntNil(nftId) {
		fileName = fmt.Sprintf(nftImageFileFormat, now.Unix())
	} else {
		fileName = fmt.Sprintf(nftImageIDFileFormat, *nftId, now.Unix())
	}

	if !model.IsStringNil(extension) {
		fileName = fmt.Sprintf("%s.%s", fileName, *extension)
	}

	return GetNftLocationPath(s.environment, fileName)
}

func (s Client) getNftAudioFilepath(nftID *int, extension *string) string {
	now := time.Now()

	var fileName string
	if model.IsIntNil(nftID) {
		fileName = fmt.Sprintf(nftAudioFileFormat, now.Unix())
	} else {
		fileName = fmt.Sprintf(nftAudioIDFileFormat, *nftID, now.Unix())
	}

	if !model.IsStringNil(extension) {
		fileName = fmt.Sprintf("%s.%s", fileName, *extension)
	}

	return GetNftLocationPath(s.environment, fileName)
}
