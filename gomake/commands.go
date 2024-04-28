package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type CommandRunner struct {
	Script_rootpath string
}

func newCommandRunner(script_rootpath string) CommandRunner {
	return CommandRunner{
		Script_rootpath: script_rootpath,
	}
}

func (cr *CommandRunner) RunLines(lines []string) {
	// chequea primero si el comando está dentro de los comandos espaciales
	// sino está lo ejecuta
	for _, command := range lines {
		command_splitted := strings.Split(command, " ")
		command_name := command_splitted[0]
		command_args := command_splitted[1:]
		if !cr.isSpecialCmdAndExecute(command_name, command_args) {
			cr.cmdRunSliceArgs(command_name, command_args)
		}
	}

}

func (cr *CommandRunner) isSpecialCmdAndExecute(cmd string, args []string) bool {
	isSpecial := false

	switch cmd {
	case "mkdir":
		isSpecial = true
		cr.cmdMkdir(args[0])
	case "cd":
		cr.cmdCd(args[0])
	}
	return isSpecial

}

func (cr *CommandRunner) cmdRunSliceArgs(name string, arg []string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("$ %s\n", cmd.String())
	return cmd.Run()
}

func (cr *CommandRunner) cmdRun(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("> %s\n", cmd.String())
	return cmd.Run()
}

func (cr *CommandRunner) cmdMkdir(newDir string) error {
	// checks is a directory exists
	// if is not then create
	newDir = cr.Script_rootpath + "/" + newDir
	_, err := os.Stat(newDir)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = cr.cmdRun("mkdir", newDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cr *CommandRunner) cmdCd(dirpath string) error {
	// checks is a directory exists
	// if is not then create
	dirpath = cr.Script_rootpath + "/" + dirpath
	_, err := os.Stat(dirpath)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = cr.cmdRun("cd", dirpath)
		if err != nil {
			return err
		}
	}

	return nil
}
