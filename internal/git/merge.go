package git

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"it.smaso/git_swiss/internal/utilities"
)

// Merge merges the source branch into the current branch
func Merge(ctx context.Context, path, source string) error {
	if !utilities.ContainsFile(path, ".git") {
		return fmt.Errorf("not executing in a git repository")
	}

	cmd := exec.Command("git", "merge", source)
	cmd.Dir = path
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to merge branch: %s", err.Error())
	}

	return nil
}
