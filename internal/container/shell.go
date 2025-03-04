package container

import (
	"fmt"
	"os"
	"os/exec"
	"rlxos/internal/color"
)

func (container *Container) ShellAt(dir string, err error) error {
	if err != nil {
		container.dumpLogs(10)
		color.Error("%s", err.Error())
	}
	cmd := exec.Command(backend, "exec", "-t", "-w", dir, "-i", container.name, "sh")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to enter shell: %v", err)
	}

	return err

}

func (container *Container) Shell(err error) error {
	return container.ShellAt("/", err)
}
