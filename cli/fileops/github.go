package fileops

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/algebananazzzzz/odyssey/cli/types"
	"gopkg.in/yaml.v3"
)

// LoadTemplateFromGitHub loads and parses a template definition from GitHub.
func LoadTemplateFromGitHub(repo string, ref string, filePath string) (*types.Template, error) {
	if repo == "" {
		return nil, errors.New("repository name cannot be empty (e.g., 'org/repo')")
	}
	if ref == "" {
		ref = "main"
	}
	if filePath == "" {
		filePath = "odyssey.template-config.yaml"
	}

	rawURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo, ref, filePath)

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 200))
		return nil, fmt.Errorf("GitHub responded with %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var tmpl types.Template
	if err := yaml.Unmarshal(data, &tmpl); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return &tmpl, nil
}
