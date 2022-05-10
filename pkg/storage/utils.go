package storage

import "fmt"

const (
	momentsLocationPath        = "%v/moments/%v/%v" // "{environment}/moments/{moment_id}/{file_name}"
	spaceLocationPath          = "%v/space/%v/%v"   // "{environment}/space/{space_id}/{file_name}"
	nftLocationPath            = "%v/nfts/%v/%v"    // "{environment}/nfts/{nft_id}/{file_name}"
	userLocationPath           = "%v/users/%v/%v"   // "{environment}/users/{user_id}/{file_name}"
	speakerProfileLocationPath = "%v/speakers/%v"   // "{environment}/speakers/{file_name}"

	publicFileFormat = "https://storage.googleapis.com/%v/%v" // https://storage.googleapis.com/BUCKET_NAME/FILE_NAME
)

func GetPublicFilePath(bucket, fileName string) string {
	return fmt.Sprintf(publicFileFormat, bucket, fileName)
}

func GetMomentsLocationPath(env, fileName string, momentID int) string {
	return fmt.Sprintf(momentsLocationPath, env, momentID, fileName)
}

func GetSpaceLocationPath(env, fileName string, spaceID int) string {
	return fmt.Sprintf(spaceLocationPath, env, spaceID, fileName)
}

func GetNftLocationPath(env, fileName string, nftID int) string {
	return fmt.Sprintf(nftLocationPath, env, nftID, fileName)
}

func GetUserLocationPath(env, fileName string, userID int) string {
	return fmt.Sprintf(userLocationPath, env, userID, fileName)
}

func GetSpeakerProfileLocationPath(env, fileName string) string {
	return fmt.Sprintf(speakerProfileLocationPath, env, fileName)
}
