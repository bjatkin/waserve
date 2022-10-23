package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(devCmd)
	rootCmd.AddCommand(serveCmd)
}

var rootCmd = &cobra.Command{
	Use:   "waserve",
	Short: "Waserve is a tool for building go wasm code",
	Long: `Waserve is a simple tool for building go wasm code
	It's primary focus is full page wasm application`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// check to make sure both npm and tinygo are installed.
		// if not give instructions for installing

		// check for npm
		if _, err := exec.LookPath("npm"); err != nil {
			return fmt.Errorf("could not find npm in path, https://docs.npmjs.com/downloading-and-installing-node-js-and-npm %s", err)
		}

		// check for tinygo
		if _, err := exec.LookPath("tinygo"); err != nil {
			return fmt.Errorf("could not find tinygo in path, https://tinygo.org/getting-started/install/ %s", err)
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
