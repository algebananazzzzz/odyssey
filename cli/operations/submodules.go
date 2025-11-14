package operations

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/algebananazzzzz/odyssey/cli/types"
)

// AddProjectSubmodules returns a task that ensures each provided submodule exists within repoPath.
func AddProjectSubmodules(repoPath string, subs []types.Submodule) func(context.Context) error {
	return func(ctx context.Context) error {
		for _, sm := range subs {
			if err := addSubmodule(ctx, repoPath, sm); err != nil {
				return err
			}
		}
		return nil
	}
}

// addSubmodule adds the given submodule to repoPath when it is missing and checks out the requested ref.
func addSubmodule(ctx context.Context, repoPath string, sm types.Submodule) error {
	if sm.Repo == "" || sm.Path == "" {
		return fmt.Errorf("invalid submodule definition: %+v", sm)
	}

	if err := runGit(ctx, repoPath, "submodule", "status", "--", sm.Path); err == nil {
		return nil
	}

	if err := runGit(ctx, repoPath, "submodule", "add", sm.Repo, sm.Path); err != nil {
		return fmt.Errorf("failed to add submodule %s: %w", sm.Path, err)
	}

	if sm.Ref == "" {
		return nil
	}

	modulePath := filepath.Join(repoPath, sm.Path)
	if err := runGit(ctx, modulePath, "checkout", sm.Ref); err != nil {
		return fmt.Errorf("failed to checkout ref %s for submodule %s: %w", sm.Ref, sm.Path, err)
	}

	return nil
}

// runGit executes a git command inside dir.
func runGit(ctx context.Context, dir string, args ...string) error {
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git %v failed: %v, %s", args, err, output)
	}
	return nil
}
