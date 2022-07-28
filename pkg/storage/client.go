package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/uncut-fm/uncut-common/pkg/errors"
	"github.com/uncut-fm/uncut-common/pkg/logger"
	"net/url"
	"strings"
	"time"
)

var (
	nftImageFileFormat   = "nft_image_%v"    // nft_image_{time_now}.{ext}
	nftAudioFileFormat   = "nft_audio_%v"    // nft_audio_{time_now}.{ext}
	nftVideoFileFormat   = "nft_video_%v"    // nft_video_{time_now}.{ext}
	nftImageIDFileFormat = "%v/nft_image_%v" // {nft_id}/nft_image_{time_now}.{ext}
	nftAudioIDFileFormat = "%v/nft_audio_%v" // {nft_id}/nft_audio_{time_now}.{ext}
	nftVideoIDFileFormat = "%v/nft_video_%v" // {nft_id}/nft_video_{time_now}.{ext}

	speakerProfileIDPath = "%v/avatar_%v.%s" // "{speaker_id}/avatar_{time_now}.{ext}"
	speakerProfilePath   = "avatar_%v.%s"    // /avatar_{time_now}.{ext}"

	spaceAttachmentFileFormat = "%v/space_attachment_%s_%v.%s" // {space_id}/space_attachment_{attachment_type}_{time_now}.{ext}

	showImageIDFileFormat = "%v/show_%v.%s" // {show_id}/show_{time_now}.{ext}
	showImageFileFormat   = "show_%v.%s"    // show_{time_now}.{ext}

	userImageIDFileFormat = "%v/user_%v.%s" // {user_id}/user_{time_now}.{ext}
	userImageFileFormat   = "user_%v.%s"    // user_{time_now}.{ext}

	audioMomentIDFileFormat = "%v/audio_%v.%s" // "{moment_id}/audio_{time_now}.{ext}"
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
	case EntityTypeSpeakerProfile:
		fileURL, err = s.storeSpeakerProfileImage(ctx, entityID, file, extension)
	case EntityTypeUser:
		fileURL, err = s.storeUserImage(ctx, entityID, file, extension)
	case EntityTypeMoment:
		fileURL, err = s.storeMomentAudio(ctx, entityID, file, extension)
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
	case EntityTypeSpeakerProfile:
		fileURL, err = s.storeSpeakerProfileImage(ctx, entityID, fileDataURLStruct.Data, extension)
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
	case EntityTypeMoment:
		fileURL, err = s.storeMomentAudio(ctx, entityID, fileDataURLStruct.Data, extension)
	}

	return fileURL, s.log.CheckError(err, s.UploadEntityFileByDataURI)
}

func (s Client) GetSignedUrl(entityType EntityType, entityID *int, mimeType string, expirationInMinutes int) (string, error) {
	extension, err := getExtensionByMimeType(mimeType)
	if s.log.CheckError(err, s.GetSignedUrl) != nil {
		return "", err
	}

	fileType := GetFileTypeByMimeType(mimeType)

	var filename string
	switch entityType {
	case EntityTypeSpace:
		filename = s.getSpaceAttachmentFilepath(*entityID, fileType, extension)
	case EntityTypeShow:
		filename = s.getShowImageFilepath(entityID, extension)
	case EntityTypeSpeakerProfile:
		filename = s.getSpeakerProfilePath(entityID, extension)
	case EntityTypeUser:
		filename = s.getUserImageFilepath(entityID, extension)
	case EntityTypeNft:
		{
			switch fileType {
			case "image":
				filename = s.getNftImageFilepath(entityID, &extension)
			case "audio":
				filename = s.getNftAudioFilepath(entityID, &extension)
			case "video":
				filename = s.getNftVideoFilepath(entityID, &extension)
			}
		}
	}

	expires := time.Now().Add(time.Minute * time.Duration(expirationInMinutes))

	signedUrl, err := s.bucketHandle.SignedURL(filename, &storage.SignedURLOptions{
		ContentType: mimeType,
		Method:      "PUT",
		Expires:     expires,
	})

	return signedUrl, s.log.CheckError(err, s.GetSignedUrl)
}

func (s Client) DeleteFileByStoragePublicURL(ctx context.Context, fileURL string) error {
	filePath, err := s.GetStorageFilePathFromPublicURL(fileURL)
	if s.log.CheckError(err, s.DeleteFileByStoragePublicURL) != nil {
		return err
	}

	err = s.deleteFileByFullFilename(ctx, filePath)
	return s.log.CheckError(err, s.DeleteFileByStoragePublicURL)
}

func (s Client) GetStorageFilePathFromPublicURL(fileURL string) (string, error) {
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

func (s Client) UploadFile(c context.Context, fileName string, file []byte) error {
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
