package operations

import (
	"context"
	"fmt"
	"time"

	"github.com/charmbracelet/huh/spinner"
)

type Task struct {
	Description string
	Action      func(context.Context) error
}

// RunTasks runs tasks sequentially, showing a spinner for each.
func RunTasks(ctx context.Context, tasks []Task) error {
	for _, task := range tasks {
		taskCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()

		err := spinner.New().
			Context(taskCtx).
			Title(task.Description).
			ActionWithErr(func(c context.Context) error {
				return task.Action(c)
			}).
			Run()
		if err != nil {
			return fmt.Errorf("task %q failed: %w", task.Description, err)
		}
	}
	return nil
}
