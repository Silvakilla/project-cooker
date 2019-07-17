package shell

import (
	"log"
	"os"
	"os/exec"
)

// Command provides a struct to do several commands in the Shell
type Command struct {
	name       string
	parameters string
}

// Shell to run shell commands
func Shell(command Command) {
	var shell = configuration.ShellToUse

	if shell == "bash" {
		cmd := exec.Command(shell, "-c", command.name+" "+command.parameters)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			log.Printf("%s", err)
		}

	} else if shell == "cmd" {
		cmd := exec.Command(shell, "/C", command.name+" "+command.parameters)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()

		if err != nil {
			log.Printf("%s", err)
		}
	}
}

func callShell(command Command) {
	Shell(command)
}

func multipleCallShell(commands []Command) {
	for _, command := range commands {
		Shell(command)
	}
}
