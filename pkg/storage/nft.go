package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
	"time"
)

func (s Client) UploadNftCsvFileByFileBytes(ctx context.Context, nftID int, file []byte, filename string) (string, error) {
	filename = s.getNftWithNameFilepath(&nftID, "csv", filename)

	// Set the content type to "text/csv"
	objectAttrs := storage.ObjectAttrs{
		ContentType: "text/csv",
		Name:        filename,
	}

	return s.uploadPublicFile(ctx, filename, file, &objectAttrs)
}

func (s Client) storeNftFile(c context.Context, nftID *int, file []byte, extension string, nftFilePath func(nftID *int, extension *string) string) (string, error) {
	fileName := nftFilePath(nftID, &extension)

	return s.uploadPublicFile(c, fileName, file, nil)
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

func (s Client) getNftWithNameFilepath(nftId *int, extension, fileName string) string {
	var filePath string

	fileName = prepareFileNameFromRequest(fileName)

	now := time.Now()
	if model.IsIntNil(nftId) {
		filePath = fmt.Sprintf(nftWithFilenameFileFormat, now.Unix(), fileName, extension)
	} else {
		filePath = fmt.Sprintf(nftWithFilenameIDFileFormat, *nftId, now.Unix(), fileName, extension)
	}

	return GetNftLocationPath(s.environment, filePath)
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
