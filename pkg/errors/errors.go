package errors

import (
	"errors"
	"fmt"
)

// twitter errors
var (
	UsernameMissingErr    = errors.New("username is missing")
	UserinfoErr           = errors.New("failed retrieving usersInfo")
	UserHandleNotFoundErr = errors.New("no user matches for specified twitter handle")
)

// show errors
var (
	NoSearchShowsFoundErr      = errors.New("showSearcher didn't found shows found with requested name")
	noSearchShowsExactFoundErr = "showSearcher didn't found shows with exactly requested name, found shows: %v"
)

func NoSearchShowsExactFoundErr(showNames []string) error {
	return fmt.Errorf(noSearchShowsExactFoundErr, showNames)
}
