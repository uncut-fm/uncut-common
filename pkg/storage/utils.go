package storage

import "fmt"

const (
	momentsLocationPath        = "%v/moments/%v/%v" // "{environment}/moments/{moment_id}/{file_name}"
	spaceLocationPath          = "%v/space/%v"      // "{environment}/space/{file_name}"
	nftLocationPath            = "%v/nfts/%v"       // "{environment}/nfts/{file_name}"
	userLocationPath           = "%v/users/%v"      // "{environment}/users/{file_name}"
	speakerProfileLocationPath = "%v/speakers/%v"   // "{environment}/speakers/{file_name}"

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
