package service

import (
	"github.com/Masterminds/semver/v3"
)

type SemverService struct {
}

func NewSemverService() SemverService {
	return SemverService{}
}

func (s SemverService) IsValid(raw string) bool {
	if _, err := semver.NewVersion(raw); err != nil {
		return false
	}

	return true
}
