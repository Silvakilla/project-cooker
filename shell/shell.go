package shell

import (
	"os"
	"os/exec"

	"project-cooker/configuration"
	"project-cooker/constants"
	fileWriter "project-cooker/file"
)

// Command provides a struct to do several commands in the Shell
type Command struct {
	name       string
	parameters string
}

// Shell to run shell commands
func Shell(command Command) {
	var shell = configuration.CurrentConfig.GetShellToUse()

	if shell == "bash" {
		cmd := exec.Command(shell, "-c", command.name+" "+command.parameters)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			fileWriter.WriteStringToFile(err.Error(), constants.ERROR)
		}

	} else if shell == "cmd" {
		cmd := exec.Command(shell, "/C", command.name+" "+command.parameters)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			fileWriter.WriteStringToFile(err.Error(), constants.ERROR)
		}
	} else {
		fileWriter.WriteStringToFile("No Shell found to use", constants.ERROR)
	}
}

// CallShell with only command to execute
func CallShell(command Command) {
	Shell(command)
}

// MultiCallShell with several commands to execute
func MultiCallShell(commands []Command) {
	for _, command := range commands {
		Shell(command)
	}
}
