package validators

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	filePathPattern = regexp.MustCompile(`^[A-Za-z0-9._-]+$`)
)

// FilePath validates that the provided directory/file path is a single
// relative segment without path separators.
func FilePath(value string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return fmt.Errorf("value is required")
	}
	if !filePathPattern.MatchString(value) {
		return fmt.Errorf("value must only include letters, numbers, dashes, underscores, or periods")
	}
	return nil
}
