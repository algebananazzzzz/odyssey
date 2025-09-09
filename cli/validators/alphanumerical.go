package validators

import (
	"errors"
	"fmt"
	"regexp"
)

func Alphanumeric(input string) error {
	if err := NotEmpty(input); err != nil {
		return err
	}

	pattern := `^[a-zA-Z0-9_][\w-]*[a-zA-Z0-9_]$`

	matched, err := regexp.MatchString(pattern, input)
	if err != nil {
		// Handle any regex errors
		fmt.Println("Regex error:", err)
		return errors.New("error in regular expression matching")
	}

	if !matched {
		return errors.New("invalid input: please only use letters, numbers, underscores, and dashes")
	}

	return nil
}
