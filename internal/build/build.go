package build

import (
	"context"
	"os"
	"os/exec"
)

func RunGoBuild(ctx context.Context) error {
	goBuild := exec.CommandContext(ctx, "tinygo", "build", "-o", "./frontend/public/main.wasm", "-target", "wasm", ".")
	goBuild.Stdout = os.Stdout
	goBuild.Stderr = os.Stderr

	return goBuild.Run()
}
