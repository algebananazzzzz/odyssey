package constants

import (
	"fmt"

	"github.com/algebananazzzzz/odyssey/cli/fileops"
	"github.com/algebananazzzzz/odyssey/cli/types"
)

var (
	TEMPLATES = make(map[string]*types.Template)
)

// LoadTemplate retrieves (and caches) the template definition for the given template key.
func LoadTemplate(templateKey string) (*types.Template, error) {
	if tmpl, ok := TEMPLATES[templateKey]; ok {
		return tmpl, nil
	}
	source, ok := PROJECT_TEMPLATES[templateKey]
	if !ok {
		return nil, fmt.Errorf("unknown template %q", templateKey)
	}
	if source.Repo == "" {
		return nil, fmt.Errorf("template %q does not declare a source repo", templateKey)
	}

	ref := source.Ref
	if ref == "" {
		ref = "main"
	}

	tmpl, err := fileops.LoadTemplateFromGitHub(source.Repo, ref, source.FilePath())
	if err != nil {
		return nil, err
	}

	if tmpl.Type == "" {
		tmpl.Type = templateKey
	}
	if tmpl.Name == "" {
		tmpl.Name = source.Name
		if tmpl.Name == "" {
			tmpl.Name = templateKey
		}
	}

	TEMPLATES[templateKey] = tmpl

	return tmpl, nil
}
