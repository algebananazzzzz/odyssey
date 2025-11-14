package validators

import (
	"fmt"
	"strings"
)

// NotEmpty validates that a string is not blank.
func NotEmpty(v string) error {
	if strings.TrimSpace(v) == "" {
		return fmt.Errorf("cannot be empty")
	}
	return nil
}
