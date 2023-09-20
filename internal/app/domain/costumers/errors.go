package costumers

import "errors"

var (
	ErrAlreadyExists      = errors.New("An user with that phone number already exists")
	ErrSearchTermTooShort = errors.New("Search term too short")
)
