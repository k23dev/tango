package main

import (
	"log"
	"os"
)

var scriptFileRootPath string = "./"
var ScriptName string
var taskName string

func main() {

	log.SetFlags(0)

	if len(os.Args) >= 2 {
		ScriptName = scriptFileRootPath + os.Args[1]
	} else {
		log.Fatalln("No script .gomake setted")
	}

	if len(os.Args) >= 3 {
		taskName = os.Args[2]
	} else {
		log.Fatalln("no task setted")
	}

	// carga el interprete
	interpreter := newInterpreter(ScriptName)

	// corre el interprete
	interpreter.Run()

	// checks if the tasks exists
	err := interpreter.CheckIfTaskExists(taskName)
	if err != nil {
		log.Fatalln(err)
	}
	// execute the task
	interpreter.RunTask()

}
