package repository

import (
	"os"
	"os/exec"
)

const DefaultRepositoryPattern = "default"

type Repository struct {
	Remote string
	Dir string
	Sha string
}

func New(remote string) *Repository {
	r := Repository{
		Remote: remote,
	}

	return &r
}

// Clone clones a repository to a tempdir and returns the dir
func (r *Repository) Clone() (string, error) {
	tmp, err := os.MkdirTemp("", DefaultRepositoryPattern)
	if err != nil {
		return "", err
	}

	command := exec.Command("git", "clone", r.Remote, tmp)
	if err := command.Run(); err != nil {
		return tmp, err
	}

	r.Dir = tmp

	return tmp, nil
}

// Checkout takes a commit sha and checkout out the repo to that commit
func (r *Repository) Checkout(sha string) error {
	command := exec.Command("git", "fetch", "origin")
	command.Dir = r.Dir
	if err := command.Run(); err != nil {
		return err
	}

	command = exec.Command("git", "checkout", sha)
	command.Dir = r.Dir
	if err := command.Run(); err != nil {
		return err
	}

	return nil
}
