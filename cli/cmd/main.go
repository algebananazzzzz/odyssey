package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/algebananazzzzz/odyssey/cli/config"
	"github.com/algebananazzzzz/odyssey/cli/ui/formatters"
	"github.com/algebananazzzzz/odyssey/cli/ui/forms"
	"github.com/algebananazzzzz/odyssey/cli/ui/styles"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
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

	form := forms.NewConfigForm(globalConfig.Config, projectConfig.Config, &confirm)
	if err := form.Run(); err != nil {
		if err == huh.ErrUserAborted {
			quit()
			return
		}
		log.Fatal(err)
	}

	if confirm {
		if err := globalConfig.Save(); err != nil {
			fmt.Printf("Failed to save global config: %v\n", err)
			os.Exit(1)
		}

		if err := projectConfig.Save(); err != nil {
			fmt.Printf("Failed to save global config: %v\n", err)
			os.Exit(1)
		}

		action := func() {
			time.Sleep(1 * time.Second)
		}
		if err := spinner.New().Title("Preparing your burger...").Action(action).Run(); err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(formatters.PrintProjectSummary(projectConfig.Config))
	} else {
		quit()
		return
	}
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
