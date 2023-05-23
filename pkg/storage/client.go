package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/uncut-fm/uncut-common/model"
	"github.com/uncut-fm/uncut-common/pkg/errors"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"net/url"
	"strings"
	"time"
)

var (
	nftImageFileFormat          = "nft_image_%v"    // nft_image_{time_now}.{ext}
	nftAudioFileFormat          = "nft_audio_%v"    // nft_audio_{time_now}.{ext}
	nftVideoFileFormat          = "nft_video_%v"    // nft_video_{time_now}.{ext}
	nftImageIDFileFormat        = "%v/nft_image_%v" // {nft_id}/nft_image_{time_now}.{ext}
	nftWithFilenameFileFormat   = "%d/%v.%v"        // {time_now}/{filename}.{ext}
	nftWithFilenameIDFileFormat = "%v/%d/%v.%v"     // {nft_id}/{time_now}/{filename}.{ext}
	nftAudioIDFileFormat        = "%v/nft_audio_%v" // {nft_id}/nft_audio_{time_now}.{ext}
	nftVideoIDFileFormat        = "%v/nft_video_%v" // {nft_id}/nft_video_{time_now}.{ext}

	conversationAttachmentWithFilenameFileFormat   = "%d/%s.%s"            // {time_now}/{filename}.{ext}
	conversationAttachmentIDWithFilenameFileFormat = "%d/%d/%s.%s"         // {conversation_id}/{time_now}/{filename}.{ext}
	conversationAttachmentIDFileFormat             = "%d/attachment_%v.%s" // {conversation_id}/attachment_{time_now}.{ext}"
	conversationAttachmentFileFormat               = "attachment_%v.%s"    // attachment_{time_now}.{ext}"

	spaceAttachmentFileFormat = "%v/space_attachment_%s_%v.%s" // {space_id}/space_attachment_{attachment_type}_{time_now}.{ext}

	showImageIDFileFormat = "%v/show_%v.%s" // {show_id}/show_{time_now}.{ext}
	showImageFileFormat   = "show_%v.%s"    // show_{time_now}.{ext}

	userImageIDFileFormat = "%v/user_%v.%s" // {user_id}/user_{time_now}.{ext}
	userImageFileFormat   = "user_%v.%s"    // user_{time_now}.{ext}
)

type Client struct {
	log                 logger.Logger
	bucketHandle        *storage.BucketHandle
	bucket, environment string
}

func NewClient(log logger.Logger, bucketHandler *storage.BucketHandle, bucket, env string) *Client {
	return &Client{
		log:          log,
		bucketHandle: bucketHandler,
		bucket:       bucket,
		environment:  env,
	}
}

func (s Client) GetBucket() string {
	return s.bucket
}

func (s Client) UploadEntityFileByFileBytes(ctx context.Context, entityType EntityType, entityID *int, file []byte, extension string) (string, error) {
	var (
		fileURL string
		err     error
	)

	switch entityType {
	case EntityTypeSpace:
		fileURL, err = s.uploadSpaceAttachmentFile(ctx, *entityID, file, extension, "image")
	case EntityTypeShow:
		fileURL, err = s.storeShowImage(ctx, entityID, file, extension)
	case EntityTypeConversation:
		fileURL, err = s.storeConversationImage(ctx, entityID, file, extension)
	case EntityTypeUser:
		fileURL, err = s.storeUserImage(ctx, entityID, file, extension)
	}

	return fileURL, s.log.CheckError(err, s.UploadEntityFileByFileBytes)
}

func (s Client) UploadEntityFileByDataURI(ctx context.Context, fileDataURLString string, entityType EntityType, entityID *int) (string, error) {
	fileDataURLStruct, err := getDataURLInfo(fileDataURLString)
	if s.log.CheckError(err, s.UploadEntityFileByDataURI) != nil {
		return "", err
	}

	var fileURL string

	extension, err := getExtensionByDataURL(fileDataURLStruct)
	if s.log.CheckError(err, s.UploadEntityFileByDataURI) != nil {
		return "", err
	}

	switch entityType {
	case EntityTypeSpace:
		fileURL, err = s.uploadSpaceAttachmentFile(ctx, *entityID, fileDataURLStruct.Data, extension, fileDataURLStruct.Type)
	case EntityTypeShow:
		fileURL, err = s.storeShowImage(ctx, entityID, fileDataURLStruct.Data, extension)
	case EntityTypeConversation:
		fileURL, err = s.storeConversationImage(ctx, entityID, fileDataURLStruct.Data, extension)
	case EntityTypeNft:
		switch fileDataURLStruct.Type {
		case "image":
			fileURL, err = s.storeNftFile(ctx, entityID, fileDataURLStruct.Data, extension, s.getNftImageFilepath)
		case "audio":
			fileURL, err = s.storeNftFile(ctx, entityID, fileDataURLStruct.Data, extension, s.getNftAudioFilepath)
		case "video":
			fileURL, err = s.storeNftFile(ctx, entityID, fileDataURLStruct.Data, extension, s.getNftVideoFilepath)
		}
	case EntityTypeUser:
		fileURL, err = s.storeUserImage(ctx, entityID, fileDataURLStruct.Data, extension)
	}

	return fileURL, s.log.CheckError(err, s.UploadEntityFileByDataURI)
}

func getExtensionAndFileType(requestedMimeType, fileName *string) (ext, fileType, mimeType string, err error) {
	if fileName == nil && requestedMimeType == nil {
		return "", "", "", errors.MimetypeErr
	}

	if model.IsStringNil(requestedMimeType) {
		ext = getExtensionFromFilename(*fileName)

		mimeType = getMimeTypeByExtension(ext)

		fileType = GetFileTypeByMimeType(mimeType)

		return
	}

	mimeType = *requestedMimeType

	ext, err = getExtensionByMimeType(mimeType)
	if err != nil {
		return
	}

	fileType = GetFileTypeByMimeType(mimeType)

	return
}

func (s Client) GetSignedUrl(entityType EntityType, entityID *int, requestedMimeType, requestedFilename *string, expirationInMinutes int) (signedURL string, mimeType string, err error) {
	extension, fileType, mimeType, err := getExtensionAndFileType(requestedMimeType, requestedFilename)
	if s.log.CheckError(err, s.GetSignedUrl) != nil {
		return "", "", err
	}

	var filename string
	switch entityType {
	case EntityTypeSpace:
		filename = s.getSpaceAttachmentFilepath(*entityID, fileType, extension)
	case EntityTypeShow:
		filename = s.getShowImageFilepath(entityID, extension)
	case EntityTypeConversation:
		if !model.IsStringNil(requestedFilename) {
			filename = s.getConversationWithNameFilepath(entityID, extension, *requestedFilename)
			break
		}

		filename = s.getConversationAttachmentFilepath(entityID, extension)
	case EntityTypeUser:
		filename = s.getUserImageFilepath(entityID, extension)
	case EntityTypeNft:
		if !model.IsStringNil(requestedFilename) {
			filename = s.getNftWithNameFilepath(entityID, extension, *requestedFilename)
			break
		}

		switch fileType {
		case "image":
			filename = s.getNftImageFilepath(entityID, &extension)
		case "audio":
			filename = s.getNftAudioFilepath(entityID, &extension)
		case "video":
			filename = s.getNftVideoFilepath(entityID, &extension)
		}
	}

	expires := time.Now().Add(time.Minute * time.Duration(expirationInMinutes))

	signedUrl, err := s.bucketHandle.SignedURL(filename, &storage.SignedURLOptions{
		ContentType: mimeType,
		Method:      "PUT",
		Expires:     expires,
	})

	return signedUrl, mimeType, s.log.CheckError(err, s.GetSignedUrl)
}

func (s Client) DeleteFileByStoragePublicURL(ctx context.Context, fileURL string) error {
	filePath, err := s.getStorageFilePathFromPublicURL(fileURL)
	if s.log.CheckError(err, s.DeleteFileByStoragePublicURL) != nil {
		return err
	}

	err = s.deleteFileByFullFilename(ctx, filePath)
	return s.log.CheckError(err, s.DeleteFileByStoragePublicURL)
}

func (s Client) getStorageFilePathFromPublicURL(fileURL string) (string, error) {
	urlParts, err := url.Parse(fileURL)
	if s.log.CheckError(err, s.DeleteFileByStoragePublicURL) != nil {
		return "", err
	}
	path := urlParts.Path
	i := strings.Index(path, s.environment)

	if i == -1 {
		return "", errors.FileAccessErr
	}
	return path[i:], nil
}

func (s Client) uploadFile(c context.Context, fileName string, file []byte) error {
	obj := s.bucketHandle.Object(fileName)

	wr := obj.NewWriter(c)

	_, err := wr.Write(file)
	if err != nil {
		return err
	}

	err = wr.Close()
	if err != nil {
		return err
	}

	return err
}

func (s Client) deleteFileByFullFilename(c context.Context, filename string) error {
	obj := s.bucketHandle.Object(filename)

	err := obj.Delete(c)
	return s.log.CheckError(err, s.deleteFileByFullFilename)
}

func (s Client) MakeFilePublic(ctx context.Context, fileName string) error {
	acl := s.bucketHandle.Object(fileName).ACL()
	err := acl.Set(ctx, storage.AllUsers, storage.RoleReader)
	if err != nil {
		return fmt.Errorf("ACLHandle.Set: %v", err)
	}

	return nil
}
