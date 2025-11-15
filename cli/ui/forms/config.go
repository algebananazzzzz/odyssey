package forms

import (
	"github.com/algebananazzzzz/odyssey/cli/types"
	"github.com/algebananazzzzz/odyssey/cli/ui/styles"
	"github.com/charmbracelet/huh"
)

// NewConfigForm builds and returns the interactive configuration wizard.
func NewConfigForm(cfg types.Config, confirm *bool) *huh.Form {
	builder := newConfigFormBuilder(cfg, confirm)
	return huh.NewForm(builder.groups()...).WithTheme(styles.NewTheme())
}

type configFormBuilder struct {
	cfg            types.Config
	confirm        *bool
	projectSelect  *huh.Select[string]
	workflowSelect *huh.Select[string]
	lastTemplate   string
}

func newConfigFormBuilder(cfg types.Config, confirm *bool) *configFormBuilder {
	if cfg.ProjectConfig == nil {
		cfg.ProjectConfig = &types.ProjectConfig{}
	}
	ensureProjectVariables(cfg.ProjectConfig)

	if cfg.ProjectConfig.Directory == "" {
		cfg.ProjectConfig.Directory = defaultProjectDirectory(cfg.ProjectConfig.Code)
	}

	builder := &configFormBuilder{
		cfg:          cfg,
		confirm:      confirm,
		lastTemplate: cfg.ProjectConfig.Template,
	}
	return builder
}

func (b *configFormBuilder) groups() []*huh.Group {
	groups := []*huh.Group{
		b.welcomeGroup(),
		b.globalSettingsGroup(),
		b.projectSettingsGroup(),
		b.templateSelectionGroup(),
		b.workflowSelectionGroup(),
	}

	if group := projectVariableGroup(b.cfg); group != nil {
		groups = append(groups, group)
	}

	groups = append(groups, b.finalConfirmationGroup())
	return groups
}

func (b *configFormBuilder) resetTemplateSelections() {
	if b.cfg.ProjectConfig == nil {
		return
	}
	b.cfg.ProjectConfig.Type = ""
	b.cfg.ProjectConfig.Workflow = ""
	if b.projectSelect != nil {
		b.projectSelect.Value(&b.cfg.ProjectConfig.Type)
	}
	if b.workflowSelect != nil {
		b.workflowSelect.Value(&b.cfg.ProjectConfig.Workflow)
	}
}
