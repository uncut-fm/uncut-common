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

	return s.uploadPublicFile(c, fileName, file, nil)
}

func (s Client) getShowImageFilepath(showID *int, extension string) string {
	now := time.Now()

	var fileName string
	if model.IsIntNil(showID) {
		fileName = fmt.Sprintf(collectionImageFileFormat, now.Unix(), extension)
	} else {
		fileName = fmt.Sprintf(collectionImageIDFileFormat, *showID, now.Unix(), extension)
	}

	return GetShowLocationPath(s.environment, fileName)
}

func (s Client) getCollectionWithNameImageFilepath(showID *int, extension, fileName string) string {
	var filePath string

	fileName = prepareFileNameFromRequest(fileName)

	now := time.Now()
	if model.IsIntNil(showID) {
		filePath = fmt.Sprintf(collectionImageWithNameFileFormat, now.Unix(), fileName, extension)
	} else {
		filePath = fmt.Sprintf(collectionImageIDWithNameFileFormat, *showID, now.Unix(), fileName, extension)
	}

	return GetShowLocationPath(s.environment, filePath)
}
