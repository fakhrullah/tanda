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

func IsGitCleanIgnoreUntracked() bool {
	// Open the repository
	repo, err := git.PlainOpen(".")
	if err != nil {
		// return false, fmt.Errorf("failed to open repo: %w", err)
		log.Error().Msgf("failed to open repo: %w", err)
		return false
	}

	// Get the working tree
	worktree, err := repo.Worktree()
	if err != nil {
		// return false, fmt.Errorf("failed to get worktree: %w", err)
		log.Error().Msgf("failed to get worktree: %w", err)
		return false
	}

	// Get the status
	status, err := worktree.Status()
	if err != nil {
		// return false, fmt.Errorf("failed to get status: %w", err)
		log.Error().Msgf("failed to get status: %w", err)
		return false
	}

	// Check if there are any staged or modified files
	// This ignores untracked files
	for _, fileStatus := range status {
		// Check if file is staged or modified (but not untracked)
		if fileStatus.Staging != git.Untracked && fileStatus.Staging != git.Unmodified {
			// return false, nil
			return false
		}
		if fileStatus.Worktree != git.Untracked && fileStatus.Worktree != git.Unmodified {
			// return false, nil
			return false
		}
	}

	return true
}

func showDirtyFiles(repoPath string) error {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	status, err := worktree.Status()
	if err != nil {
		return err
	}

	fmt.Println("Files with uncommitted changes:")
	for file, fileStatus := range status {
		if fileStatus.Staging != git.Untracked && fileStatus.Staging != git.Unmodified {
			fmt.Printf("  %s (staged: %s)\n", file, fileStatus.Staging)
		}
		if fileStatus.Worktree != git.Untracked && fileStatus.Worktree != git.Unmodified {
			fmt.Printf("  %s (worktree: %s)\n", file, fileStatus.Worktree)
		}
	}

	return nil
}
