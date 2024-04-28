// task.go
//go:build ignore

package main

import (
	"log"
	"os"
	"os/exec"
)

const BINARY_NAME = "cli"
const BINARY_NAME_WIN = "cli.exe"

const BUILD_DIR = "../build"

const BUILD_DIR_LINUX = "../build/linux"
const BUILD_DIR_LINUXARM64 = "../build/arm64"
const BUILD_DIR_WIN = "../build/windows"

func cmdRun(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("$ %s\n", cmd.String())
	return cmd.Run()
}

func main() {
	log.SetFlags(0)
	var taskName string
	if len(os.Args) >= 2 {
		taskName = os.Args[1]
	} else {
		log.Fatalln("no task")
	}
	task, ok := map[string]func() error{
		"install": Install,
		"build":   Build,
		"test":    Test,
		// Add more tasks here!
	}[taskName]
	if !ok {
		log.Fatalln("no such task")
	}
	err := task()
	if err != nil {
		log.Fatalln(err)
	}
}

// Tasks comes here

func Build() error {
	err := cmdRun("go", "build", "-o", ".out/", "-tags", "embed,nonet,purego", "./cmd/tool-one")
	if err != nil {
		return err
	}
	err = cmdRun("go", "build", "-o", ".out/", "-tags", "octokit,sqlite", "./cmd/tool-two")
	if err != nil {
		return err
	}
	// ...
	return nil
}

func Install() error {

	err := cmdRun("go", "mod", "tidy")
	if err != nil {
		return err
	}

	return nil
}

func Test() error {
	err := cmdRun("go", "test", "./tests")
	if err != nil {
		return err
	}

	return nil
}
