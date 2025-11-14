package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/algebananazzzzz/odyssey/cli/config"
	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/operations"
	"github.com/algebananazzzzz/odyssey/cli/types"
	"github.com/algebananazzzzz/odyssey/cli/ui/formatters"
	"github.com/algebananazzzzz/odyssey/cli/ui/forms"
	"github.com/algebananazzzzz/odyssey/cli/ui/styles"
	"github.com/charmbracelet/huh"
)

// --- Main ---
// Execute runs the CLI flow, collects configuration, and orchestrates scaffolding tasks.
func Execute() {
	var confirm bool = true
	globalConfig, err := config.LoadGlobalConfig()
	if err != nil {
		fmt.Printf("Failed to load global config: %v\n", err)
		os.Exit(1)
	}
	projectConfig, err := config.LoadProjectConfig()
	if err != nil {
		var missing config.ErrMissingProjectVariable
		if errors.As(err, &missing) {
			var reconfigure bool
			if !runDecisionForm(forms.NewMissingVariablePrompt(missing, &reconfigure), &reconfigure) {
				quit()
				return
			}
		} else {
			fmt.Printf("Failed to load project config: %v\n", err)
			os.Exit(1)
		}
	} else if projectConfig.Corrupted {
		var reconfigure bool
		if !runDecisionForm(forms.NewCorruptionPrompt(&reconfigure), &reconfigure) {
			quit()
			return
		}
	}
	cfg := types.Config{
		GlobalConfig:  globalConfig.Config,
		ProjectConfig: projectConfig.Config,
	}

	form := forms.NewConfigForm(cfg, &confirm)
	if err := form.Run(); err != nil {
		if err == huh.ErrUserAborted {
			quit()
			return
		}
		log.Fatal(err)
	}

	if !confirm {
		quit()
		return
	}

	if err := globalConfig.Save(); err != nil {
		fmt.Printf("Failed to save global config: %v\n", err)
		os.Exit(1)
	}

	if err := projectConfig.Save(); err != nil {
		fmt.Printf("Failed to save global config: %v\n", err)
		os.Exit(1)
	}

	targetDir := cfg.ProjectConfig.Directory
	if targetDir == "" {
		targetDir = "."
	}
	if !filepath.IsAbs(targetDir) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get current working directory: %v", err)
		}
		targetDir = filepath.Join(cwd, targetDir)
	}

	tmpl, err := constants.LoadTemplate(projectConfig.Config.Template)
	if err != nil {
		log.Fatalf("Failed to load template metadata: %v", err)
	}
	projectDef, err := tmpl.Project(projectConfig.Config.Type)
	if err != nil {
		log.Fatalf("Failed to resolve project %q: %v", projectConfig.Config.Type, err)
	}

	submodules := projectDef.Submodules

	tasks := []operations.Task{
		{
			Description: "Creating project directory...",
			Action:      operations.CreateProjectDirectory(targetDir),
		},
		{
			Description: "Cloning project files...",
			Action:      operations.CloneProjectFiles(targetDir, cfg),
		},
		{
			Description: "Initializing Git repository...",
			Action:      operations.InitGit(targetDir),
		},
		{
			Description: "Adding project submodules...",
			Action:      operations.AddProjectSubmodules(targetDir, submodules),
		},
		// {
		// 	Description: "Customizing content files...",
		// 	Action:      operations.CustomizeContentFiles(targetDir, cfg),
		// },
	}

	ctx := context.Background()

	if err := operations.RunTasks(ctx, tasks); err != nil {
		log.Fatalf("One or more tasks failed: %v", err)
	}

	fmt.Println(formatters.PrintProjectSummary(projectConfig.Config))
}

// quit clears the screen and prints the farewell message.
func quit() {
	fmt.Println("\033[H\033[2J")
	fmt.Println(styles.ExitStyle.Render("OdysseyCli - See you next time! 👋"))
}

func runDecisionForm(form *huh.Form, decision *bool) bool {
	if err := form.Run(); err != nil {
		if err == huh.ErrUserAborted {
			return false
		}
		log.Fatal(err)
	}
	return *decision
}
