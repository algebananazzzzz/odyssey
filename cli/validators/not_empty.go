package validators

import (
	"fmt"
	"strings"
)

func NotEmpty(v string) error {
	if strings.TrimSpace(v) == "" {
		return fmt.Errorf("cannot be empty")
	}
	return nil
}
