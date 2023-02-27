package storage

import (
	"fmt"
	"github.com/vincent-petithory/dataurl"
	"mime"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	momentsLocationPath        = "%v/moments/%v/%v" // "{environment}/moments/{moment_id}/{file_name}"
	spaceLocationPath          = "%v/spaces/%v"     // "{environment}/spaces/{file_name}"
	nftLocationPath            = "%v/nfts/%v"       // "{environment}/nfts/{file_name}"
	userLocationPath           = "%v/users/%v"      // "{environment}/users/{file_name}"
	speakerProfileLocationPath = "%v/speakers/%v"   // "{environment}/speakers/{file_name}"
	showLocationPath           = "%v/shows/%v"      // "{environment}/shows/{file_name}"

	publicFileFormat = "https://storage.googleapis.com/%v/%v" // https://storage.googleapis.com/BUCKET_NAME/FILE_NAME
)

func GetPublicFilePath(bucket, fileName string) string {
	return fmt.Sprintf(publicFileFormat, bucket, fileName)
}

func GetMomentsLocationPath(env, fileName string, momentID int) string {
	return fmt.Sprintf(momentsLocationPath, env, momentID, fileName)
}

func GetSpaceLocationPath(env, fileName string) string {
	return fmt.Sprintf(spaceLocationPath, env, fileName)
}

func GetNftLocationPath(env, fileName string) string {
	return fmt.Sprintf(nftLocationPath, env, fileName)
}

func GetUserLocationPath(env, fileName string) string {
	return fmt.Sprintf(userLocationPath, env, fileName)
}

func GetSpeakerProfileLocationPath(env, fileName string) string {
	return fmt.Sprintf(speakerProfileLocationPath, env, fileName)
}

func GetShowLocationPath(env, fileName string) string {
	return fmt.Sprintf(showLocationPath, env, fileName)
}

var mimeTypesToExt = map[string]string{
	"audio/mpeg":         "mp3",
	"image/png":          "png",
	"image/jpeg":         "jpeg",
	"image/gif":          "gif",
	"video/mp4":          "mp4",
	"video/x-flv":        "flv",
	"video/octet-stream": "mkv",
	"video/3gpp":         "3gp",
	"video/webm":         "webm",
	"video/quicktime":    "mov",
}

func getExtensionByMimeType(mimeType string) (string, error) {
	if ext, ok := mimeTypesToExt[mimeType]; ok {
		return ext, nil
	}

	extensions, err := mime.ExtensionsByType(mimeType)
	if err != nil {
		return "", err
	}

	if len(extensions) == 0 {
		mimetypeParts := strings.SplitN(mimeType, "/", 2)
		return mimetypeParts[1], nil
	}

	// return without leading dot
	return extensions[0][1:], nil
}

// accepts extension without leading dot
func getMimeTypeByExtension(extension string) string {
	return mime.TypeByExtension(extension)
}

// getDataURLInfo parses dataURL string and retrieves bytes
func getDataURLInfo(dataURLString string) (*dataurl.DataURL, error) {
	return dataurl.DecodeString(dataURLString)
}

func getExtensionByDataURL(data *dataurl.DataURL) (string, error) {
	mimeType := fmt.Sprintf("%s/%s", data.Type, data.Subtype)

	return getExtensionByMimeType(mimeType)
}

func GetFileTypeByMimeType(mimeType string) string {
	parts := strings.SplitN(mimeType, "/", 2)
	return parts[0]
}

func prepareFileNameFromRequest(filename string) string {
	filename = strings.Replace(filename, " ", "_", -1)
	return filepath.Base(filename)
}

func getExtensionFromFilename(filename string) string {
	return filepath.Ext(filename)
}

func GetEntityTypeAndEntityIDByObjectName(objectName string) (EntityType, int, error) {
	parts := strings.Split(objectName, "/")
	if len(parts) < 3 {
		return "", 0, nil
	}

	entityID, err := strconv.Atoi(parts[2])
	if err != nil {
		return EntityTypeShow, 0, err
	}

	var entityType EntityType

	switch parts[1] {
	case "nfts":
		entityType = EntityTypeNft
	case "shows":
		entityType = EntityTypeShow
	case "spaces":
		entityType = EntityTypeSpace
	case "users":
		entityType = EntityTypeUser
	case "speakers":
		entityType = EntityTypeSpeakerProfile
	}

	return entityType, entityID, nil
}
