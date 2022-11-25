package service

import (
	"github.com/Masterminds/semver/v3"
)

type SemverUp int

const (
	SemverUpMajor SemverUp = iota
	SemverUpMinor
	SemverUpPatch
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

func (s SemverService) Up(raw string, up SemverUp) string {
	v, err := semver.NewVersion(raw)
	if err != nil {
		v, _ = semver.NewVersion("0.0.0")
	}

	if up == SemverUpMajor {
		return v.IncMajor().String()
	} else if up == SemverUpMinor {
		return v.IncMinor().String()
	} else if up == SemverUpPatch {
		return v.IncPatch().String()
	}

	return raw
}
