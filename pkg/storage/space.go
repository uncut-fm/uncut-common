package storage

import (
	"context"
	"fmt"

	"time"
)

func (s Client) uploadSpaceAttachmentFile(c context.Context, spaceID int, file []byte, extension string, mediaType string) (string, error) {
	attachmentURL, err := s.storeSpaceAttachment(c, spaceID, file, mediaType, extension)

	return attachmentURL, s.log.CheckError(err, s.uploadSpaceAttachmentFile)
}

func (s Client) storeSpaceAttachment(c context.Context, spaceID int, file []byte, mediaType string, extension string) (string, error) {
	fileName := s.getSpaceAttachmentFilepath(spaceID, mediaType, extension)

	return s.uploadPublicFile(c, fileName, file, nil)
}

func (s Client) getSpaceAttachmentFilepath(spaceID int, mediaType string, extension string) string {
	now := time.Now()
	fileName := fmt.Sprintf(spaceAttachmentFileFormat, spaceID, mediaType, now.Unix(), extension)

	return GetSpaceLocationPath(s.environment, fileName)
}
