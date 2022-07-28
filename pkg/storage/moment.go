package storage

import (
	"context"
	"fmt"
	"time"
)

func (s Client) storeMomentAudio(c context.Context, userID *int, file []byte, extension string) (string, error) {
	fileName := s.getUserImageFilepath(userID, extension)

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

func (s Client) geMomentAudioFilepath(momentID int, extension string) string {
	now := time.Now()

	var fileName string

	fileName = fmt.Sprintf(audioMomentIDFileFormat, momentID, now.Unix(), extension)

	return GetMomentsLocationPath(s.environment, fileName)
}
