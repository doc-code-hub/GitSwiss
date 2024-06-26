package git

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"it.smaso/git_swiss/internal/utilities"
)

// Checkout checks out the given branch
func Checkout(ctx context.Context, path, branch string) error {
	if !utilities.ContainsFile(path, ".git") {
		return fmt.Errorf("not executing in a git repository")
	}

	if uncommitted, err := PendingChanges(ctx, path); err != nil {
		return err
	} else if uncommitted {
		return fmt.Errorf("uncommitted files in the repository")
	}

	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = path
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to checkout branch: %s", err.Error())
	}

	return nil
}

// CheckoutCreate creates a new branch from the current one and pushes it to repository
func CheckoutCreate(ctx context.Context, path, branch string) error {
	if !utilities.ContainsFile(path, ".git") {
		return fmt.Errorf("not executing in a git repository")
	}

	if uncommitted, err := PendingChanges(ctx, path); err != nil {
		return err
	} else if uncommitted {
		return fmt.Errorf("uncommitted files in the repository")
	}

	cmd := exec.Command("git", "checkout", "-b", branch)
	cmd.Dir = path
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to checkout branch: %s", err.Error())
	}

	cmd = exec.Command("git", "push", "-u", "origin")
	cmd.Dir = path
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to push new branch to origin: %s", err.Error())
	}

	return nil
}
