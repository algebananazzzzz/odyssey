package validators

import (
	"fmt"

	"github.com/algebananazzzzz/odyssey/cli/constants"
)

func AWSRegion(input string) error {
	if _, ok := constants.AWS_REGIONS[input]; !ok {
		return fmt.Errorf("invalid AWS region: %s", input)
	}
	return nil
}
