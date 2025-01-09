package repo

import "errors"

var (
	ErrAuthorNotFound   = errors.New("author not found")
	ErrGenreNotFound    = errors.New("genre not found")
	ErrLanguageNotFound = errors.New("language not found")
	ErrBookNotFound     = errors.New("book not found")
)
