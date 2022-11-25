package service

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type GitService struct {
	Repo *git.Repository
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
