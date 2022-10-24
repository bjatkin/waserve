package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bjatkin/waserve/internal/build"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build your go wasm and react code",
	RunE: func(cmd *cobra.Command, args []string) error {
		// check for a frontend dir
		info, err := os.Stat("./frontend")
		if err != nil {
			return fmt.Errorf("could not find ./frontend dir, %s", err)
		}
		if !info.IsDir() {
			return fmt.Errorf("./frontend is not a directory, %s", err)
		}

		// check for a go.mod file
		if _, err := os.Stat("go.mod"); err != nil {
			return fmt.Errorf("could not find go.mod in working dir %s", err)
		}

		// build front end code
		npmBuild := exec.CommandContext(cmd.Context(), "npm", "run", "build")
		npmBuild.Dir = "./frontend"
		npmBuild.Stdout = os.Stdout
		npmBuild.Stderr = os.Stderr

		// start this running and then move on since it can take a second
		err = npmBuild.Start()
		if err != nil {
			return fmt.Errorf("failed to start npm build %s", err)
		}

		// run go build
		if err := build.RunGoBuild(cmd.Context()); err != nil {
			return fmt.Errorf("go build failed %s", err)
		}

		// wait for both commands to finish
		err = npmBuild.Wait()
		if err != nil {
			return fmt.Errorf("npm build failed %s", err)
		}

		return nil
	},
}
