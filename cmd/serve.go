package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start a local server",
	Long:  "starts a local server using npm, it servers content from the ./dist directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		serveCmd := exec.CommandContext(cmd.Context(), "npm", "run", "preview")
		serveCmd.Dir = "./frontend"
		serveCmd.Stdout = os.Stdout
		serveCmd.Stderr = os.Stderr

		err := serveCmd.Run()
		if err != nil {
			return fmt.Errorf("failed to start preview server %s", err)
		}

		return nil
	},
}
