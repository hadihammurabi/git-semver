package service

import (
	"fmt"
	"os"
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/hadihammurabi/go-ioc/ioc"
)

type GitService struct {
	Repo          *git.Repository
	SemverService *SemverService
}

func NewGitService() GitService {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repo, err := git.PlainOpen(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return GitService{
		repo,
		ioc.Get(SemverService{}),
	}
}

func (s GitService) GetTags() []string {
	result := make([]string, 0)

	tags, err := s.Repo.Tags()
	if err != nil {
		return result
	}

	defer tags.Close()

	tags.ForEach(func(r *plumbing.Reference) error {
		result = append(result, r.Name().Short())
		return nil
	})

	return result
}

func (s GitService) GetValidTags() []string {
	tags := s.GetTags()

	vs := make([]*semver.Version, 0)
	for _, tag := range tags {
		v, err := semver.NewVersion(tag)
		if err != nil {
			continue
		}

		vs = append(vs, v)
	}

	sort.Sort(semver.Collection(vs))

	result := make([]string, len(vs))
	for _, v := range vs {
		result = append(result, v.Original())
	}
	return result
}
