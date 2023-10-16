package storage

import (
	"context"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
	"time"
)

func (s Client) StoreConversationImage(c context.Context, conversationID *int, file []byte, extension string) (string, error) {
	return s.storeConversationImage(c, conversationID, file, extension)
}

func (s Client) storeConversationImage(c context.Context, conversationID *int, file []byte, extension string) (string, error) {
	fileName := s.getConversationAttachmentFilepath(conversationID, extension)

	return s.uploadPublicFile(c, fileName, file, nil)
}

func (s Client) getConversationWithNameFilepath(conversationID *int, extension, fileName string) string {
	var filePath string

	fileName = prepareFileNameFromRequest(fileName)

	now := time.Now()

	if model.IsIntNil(conversationID) {
		filePath = fmt.Sprintf(conversationAttachmentWithFilenameFileFormat, now.Unix(), fileName, extension)
	} else {
		filePath = fmt.Sprintf(conversationAttachmentIDWithFilenameFileFormat, *conversationID, now.Unix(), fileName, extension)
	}

	return GetConversationLocationPath(s.environment, filePath)
}

func (s Client) getConversationAttachmentFilepath(conversationID *int, extension string) string {
	now := time.Now()

	var fileName string
	if model.IsIntNil(conversationID) {
		fileName = fmt.Sprintf(conversationAttachmentFileFormat, now.Unix(), extension)
	} else {
		fileName = fmt.Sprintf(conversationAttachmentIDFileFormat, *conversationID, now.Unix(), extension)
	}

	return GetConversationLocationPath(s.environment, fileName)
}
