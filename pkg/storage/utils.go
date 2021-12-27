package storage

import "fmt"

const (
	momentsLocationPath = "%v/moments/%v/%v"                     // "{environment}/moments/{moment_id}/{file_name}"
	publicFileFormat    = "https://storage.googleapis.com/%v/%v" // https://storage.googleapis.com/BUCKET_NAME/FILE_NAME
)

func GetPublicFilePath(bucket, fileName string) string {
	return fmt.Sprintf(publicFileFormat, bucket, fileName)
}

func GetMomentsLocationPath(env, fileName string, momentID int) string {
	return fmt.Sprintf(momentsLocationPath, env, momentID, fileName)
}
