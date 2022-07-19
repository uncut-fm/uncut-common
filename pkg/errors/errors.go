package errors

import (
	"errors"
	"fmt"
)

// storage errors
var (
	FileAccessErr = errors.New("requested file cannot be deleted")
)

// twitter errors
var (
	UsernameMissingErr    = errors.New("username is missing")
	UserinfoErr           = errors.New("failed retrieving usersInfo")
	UserHandleNotFoundErr = errors.New("no user matches for specified twitter handle")
)

// show errors
var (
	ShowDoesntExistErr         = errors.New("requested show doesn't exist")
	NoSearchShowsFoundErr      = errors.New("showSearcher didn't found shows found with requested name")
	noSearchShowsExactFoundErr = "showSearcher didn't found shows with exactly requested name, found shows: %v"
	ShowEpisodeDoesntExistsErr = errors.New("episode is not linked to any show")
	ShowMomentDoesntExistsErr  = errors.New("moment is not linked to any show")
)

// nft errors
var (
	NftRelationTypeErr     = errors.New("requested relationType is not supported")
	NftRelationEntityIDErr = errors.New("entityID with requested relationType must be provided")
)

// episode errors
var (
	EpisodeDoesntExistErr = errors.New("requested episode doesn't exist")
)

// moment errors
var (
	MomentDoesntExistErr = errors.New("requested moment doesn't exist")
)

// collection errors
var (
	CollectionEpisodeDoesntExistsErr = errors.New("nftCollection doesn't exist in requested episode")
	CollectionMomentDoesntExistsErr  = errors.New("nftCollection doesn't exist in requested moment")
	CollectionShowDoesntExistsErr    = errors.New("nftCollection doesn't exist in requested show")
)

func NoSearchShowsExactFoundErr(showNames []string) error {
	return fmt.Errorf(noSearchShowsExactFoundErr, showNames)
}
