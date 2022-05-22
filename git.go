package tanda

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rs/zerolog/log"
)

func AddGitCommit(version string) (plumbing.Hash, error) {
	log.Debug().Msgf("Should create commit with title version : %s", version)

	r, err := git.PlainOpen(".")
	if err != nil {
		return plumbing.ZeroHash, err
	}

	w, err := r.Worktree()
	if err != nil {
		return plumbing.ZeroHash, err
	}

	addHash, err := w.Add("pubspec.yaml")
	if err != nil {
		return plumbing.ZeroHash, err
	}
	log.Debug().Msgf("Add hash: %v", addHash)

	commitHash, err := w.Commit(version, &git.CommitOptions{})
	if err != nil {
		return plumbing.ZeroHash, err
	}
	log.Debug().Msgf("Commit hash: %v", commitHash)
	return commitHash, nil
}

func IsGitClean() bool {
	r, err := git.PlainOpen(".")
	if err != nil {
		// log err
		return false
	}

	w, err := r.Worktree()
	if err != nil {
		// log err
		return false
	}

	status, err := w.Status()
	if err != nil {
		// log err
		return false
	}

	return status.IsClean()
}

func AddGitTag(semverVersion string, hash plumbing.Hash) {
	r, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}

	tagname := fmt.Sprintf("v%v", semverVersion)
	tagRef, err := r.CreateTag(tagname, hash, &git.CreateTagOptions{
		Message: tagname,
	})
	if err != nil {
		panic(err)
	}
	log.Debug().Msgf("TagRef: %v", tagRef)
}
