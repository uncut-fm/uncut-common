package storage

import (
	"context"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
	"time"
)

func (s Client) StoreShowImage(c context.Context, showID *int, file []byte, extension string) (string, error) {
	return s.storeShowImage(c, showID, file, extension)
}

func (s Client) storeShowImage(c context.Context, showID *int, file []byte, extension string) (string, error) {
	fileName := s.getShowImageFilepath(showID, extension)

	err := s.UploadFile(c, fileName, file)
	if err != nil {
		return "", err
	}

	err = s.MakeFilePublic(c, fileName)
	if err != nil {
		return "", err
	}

	return GetPublicFilePath(s.bucket, fileName), nil
}

func (s Client) getShowImageFilepath(showID *int, extension string) string {
	now := time.Now()

	var fileName string
	if model.IsIntNil(showID) {
		fileName = fmt.Sprintf(showImageFileFormat, now.Unix(), extension)
	} else {
		fileName = fmt.Sprintf(showImageIDFileFormat, *showID, now.Unix(), extension)
	}

	return GetShowLocationPath(s.environment, fileName)
}
