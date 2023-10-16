package storage

import (
	"context"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
	"time"
)

func (s Client) StoreUserImage(c context.Context, userID *int, file []byte, extension string) (string, error) {
	return s.storeUserImage(c, userID, file, extension)
}

func (s Client) storeUserImage(c context.Context, userID *int, file []byte, extension string) (string, error) {
	fileName := s.getUserImageFilepath(userID, extension)

	return s.uploadPublicFile(c, fileName, file, nil)
}

func (s Client) getUserImageFilepath(userID *int, extension string) string {
	now := time.Now()

	var fileName string
	if model.IsIntNil(userID) {
		fileName = fmt.Sprintf(userImageFileFormat, now.Unix(), extension)
	} else {
		fileName = fmt.Sprintf(userImageIDFileFormat, *userID, now.Unix(), extension)
	}

	return GetUserLocationPath(s.environment, fileName)
}
