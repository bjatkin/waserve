package cmd

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed frontend
var frontendFS embed.FS

//go:embed go
var goFS embed.FS

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init a new waserve project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("writing frontend files...")
		err := copyDirAll(frontendFS, "", "")
		if err != nil {
			return fmt.Errorf("failed to copy frontend files to the dest %s", err)
		}

		// run npm install so things are ready to go
		npmCmd := exec.CommandContext(cmd.Context(), "npm", "install")
		npmCmd.Dir = "./frontend"
		err = npmCmd.Start()
		if err != nil {
			return fmt.Errorf("failed to start npm install command %s", err)
		}

		fmt.Println("writing go files...")
		err = copyDirAll(goFS, "go", ".tmp")
		if err != nil {
			return fmt.Errorf("failed to copy go files to the dest %s", err)
		}

		fmt.Println("waiting on npm install...")
		err = npmCmd.Wait()
		if err != nil {
			return fmt.Errorf("failed to complete npm install command %s", err)
		}

		return nil
	},
}

func copyDirAll(efs embed.FS, stripPrefix, stripSuffix string) error {
	err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == "." {
			// nothing to do if this is the root
			return nil
		}

		destPath := strings.TrimPrefix(path, stripPrefix)
		destPath = strings.TrimPrefix(destPath, "/") // make sure there's not a leading /
		destPath = strings.TrimSuffix(destPath, stripSuffix)
		if len(destPath) == 0 {
			// nothing to do in this case
			return nil
		}

		if d.IsDir() {
			err := os.MkdirAll(destPath, 0755)
			if err != nil {
				return fmt.Errorf("failed to make dir %s, %s", d.Name(), err)
			}
			return nil
		}

		fileData, err := efs.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s, %s", path, err)
		}

		err = os.WriteFile(destPath, fileData, 0o0666)
		if err != nil {
			return fmt.Errorf("failed to write file %s, %s", path, err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to copy dir %s", err)
	}

	return nil
}
