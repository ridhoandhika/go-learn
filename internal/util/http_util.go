package util

import (
	"errors"
	"ridhoandhika/backend-api/domain"
)

func GetHttpStatus(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	default:
		return 500
	}
}
