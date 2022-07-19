package storage

import (
	"context"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
	"time"
)

func (s Client) StoreSpeakerProfileImage(c context.Context, speakerID *int, file []byte, extension string) (string, error) {
	return s.storeSpeakerProfileImage(c, speakerID, file, extension)
}

func (s Client) storeSpeakerProfileImage(c context.Context, speakerID *int, file []byte, extension string) (string, error) {
	fileName := s.getSpeakerProfilePath(speakerID, extension)

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

func (s Client) getSpeakerProfilePath(speakerID *int, extension string) string {
	now := time.Now()

	var fileName string
	if model.IsIntNil(speakerID) {
		fileName = fmt.Sprintf(speakerProfilePath, now.Unix(), extension)
	} else {
		fileName = fmt.Sprintf(speakerProfileIDPath, *speakerID, now.Unix(), extension)
	}

	return GetSpeakerProfileLocationPath(s.environment, fileName)
}
