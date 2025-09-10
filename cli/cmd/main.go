package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

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
func Execute() {
	var confirm bool = true
	globalConfig, err := config.LoadGlobalConfig()
	if err != nil {
		fmt.Printf("Failed to load global config: %v\n", err)
		os.Exit(1)
	}
	projectConfig, err := config.LoadProjectConfig()
	if err != nil {
		fmt.Printf("Failed to load project config: %v\n", err)
		os.Exit(1)
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

	// Determine current directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	tasks := []operations.Task{
		{
			Description: "Cloning project files...",
			Action:      operations.CloneProjectFiles(currentDir, *projectConfig.Config),
		},
		{
			Description: "Initializing Git repository...",
			Action:      operations.InitGit(currentDir),
		},
		{
			Description: "Adding submodule...",
			Action:      operations.AddSubmodule(currentDir, "infra/modules", constants.TERRAFORM_SUBMODULE_GIT_URL),
		},
		{
			Description: "Customizing content files...",
			Action:      operations.CustomizeContentFiles(currentDir, cfg),
		},
	}

	ctx := context.Background()

	if err := operations.RunTasks(ctx, tasks); err != nil {
		log.Fatalf("One or more tasks failed: %v", err)
	}

	fmt.Println(formatters.PrintProjectSummary(projectConfig.Config))
}

func quit() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	default: // unix-like
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	fmt.Println(styles.ExitStyle.Render("OdysseyCli - See you next time! 👋"))
}
