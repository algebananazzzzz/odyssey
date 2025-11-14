package operations

import (
	"context"
	"fmt"
	"os"
)

// CreateProjectDirectory makes sure the target directory exists before scaffolding files.
func CreateProjectDirectory(path string) func(context.Context) error {
	return func(ctx context.Context) error {
		info, err := os.Stat(path)
		if err == nil {
			if !info.IsDir() {
				return fmt.Errorf("%s exists but is not a directory", path)
			}
			return nil
		}
		if os.IsNotExist(err) {
			if mkErr := os.MkdirAll(path, 0o755); mkErr != nil {
				return fmt.Errorf("failed to create project directory %s: %w", path, mkErr)
			}
			return nil
		}
		return fmt.Errorf("failed to inspect project directory %s: %w", path, err)
	}
}
