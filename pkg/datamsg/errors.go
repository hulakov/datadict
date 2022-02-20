package datamsg

import "errors"

var (
	ErrBadJson = errors.New("cannot parse JSON")
	ErrBadKind = errors.New("cannot parse message kind")
)
