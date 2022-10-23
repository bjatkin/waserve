package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bjatkin/waserve/internal/build"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "start a dev server",
	Long: `starts a dev server and listens for updates
	updates to *.go files will recomplie main.wasm file
	updates to the front end app will reload the dev server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// start the npm dev server
		npmCmd := exec.CommandContext(cmd.Context(), "npm", "run", "dev")
		npmCmd.Dir = "./frontend"
		npmCmd.Stdout = os.Stdout
		npmCmd.Stderr = os.Stderr

		err := npmCmd.Start()
		if err != nil {
			return fmt.Errorf("failed to start frontend server %s", err)
		}

		// watch for go file updates and rebuild the main.wasm
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return fmt.Errorf("failed to create new fsnotify watcher %s", err)
		}
		defer watcher.Close()

		done := make(chan bool)
		go func() {
			defer close(done)

			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if filepath.Ext(event.Name) != ".go" {
						return
					}
					if event.Op.Has(fsnotify.Create) || event.Op.Has(fsnotify.Write) || event.Op.Has(fsnotify.Remove) {
						fmt.Println("file change detected rebuilding wasm.main")
						err := build.RunGoBuild(cmd.Context())
						if err != nil {
							fmt.Printf("go build failed %s\n", err)
						}
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					fmt.Printf("filewatch error: %s\n", err)
				}
			}

		}()

		err = watcher.Add("./")
		if err != nil {
			return fmt.Errorf("failed to add watcher directory %s", err)
		}
		<-done

		return nil
	},
}
