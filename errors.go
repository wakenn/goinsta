package goinsta

import "errors"

// ErrNotFound is returned if the request responds with a 404 status code
// i.e a non existent user
var ErrNotFound = errors.New("The specified data wasn't found.")

// ErrLoggedOut is returned if the request responds with a 400 status code
var ErrLoggedOut = errors.New("The account is logged out")

var ErrBadPassword = errors.New("The password provided is incorrect")

var ErrChallenge = errors.New("Challenge required by Instagram")
