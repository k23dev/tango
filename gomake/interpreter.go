package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

const script_extension = ".gomake"

type Interpreter struct {
	currentLine        int
	filepath           string
	workingfile        *os.File
	Script_rootpath    string
	Script_variables   map[string]string
	Script_tasks       []string
	Script_taskLines   map[string][]string
	isWorkingOnTask    bool
	workingCurrentTask string
	task2run           string
}

func newInterpreter(filepath string) Interpreter {
	return Interpreter{
		currentLine:        0,
		filepath:           filepath + script_extension,
		Script_rootpath:    "",
		Script_variables:   make(map[string]string),
		Script_tasks:       []string{},
		Script_taskLines:   make(map[string][]string),
		isWorkingOnTask:    false,
		workingCurrentTask: "",
	}
}

func (in *Interpreter) Run() {

	in.OpenFile()
	in.Scann()
}

func (in *Interpreter) OpenFile() {
	// Abre el archivo
	file, err := os.Open(in.filepath)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	in.workingfile = file
}

func (in *Interpreter) Scann() {
	defer in.workingfile.Close()
	scanner := bufio.NewScanner(in.workingfile)

	// Escanea el archivo línea por línea
	for scanner.Scan() {
		line := scanner.Text()
		in.WorkLine(line)
		// Comprueba si hubo un error durante el escaneo del archivo
		if err := scanner.Err(); err != nil {
			fmt.Println("Error al escanear el archivo:", err)
			return
		}
	}

}

func (in *Interpreter) WorkLine(line string) {
	in.currentLine++
	// Define las expresiones regulares para verificar si la línea comienza con $ o //
	commentRegex := regexp.MustCompile(`^\s*\/\/`)
	dollarRegex := regexp.MustCompile(`^\$`)
	letterRegex := regexp.MustCompile(`^[a-zA-Z]`)
	endBracketRegex := regexp.MustCompile(`^\}`)
	indentRegex := regexp.MustCompile(`^[\t ]+`)

	if commentRegex.MatchString(line) {
		// fmt.Println(in.currentLine, " > Comentario")
		return
	}

	if dollarRegex.MatchString(line) {
		// fmt.Println("> Variable")
		in.workLineAsVariable(line)
	} else if letterRegex.MatchString(line) {
		// fmt.Println(in.currentLine, " > Tarea")
		in.isWorkingOnTask = true
		in.workLineAsTask(line)
	} else if indentRegex.MatchString(line) && in.isWorkingOnTask {
		// fmt.Println(in.currentLine, " >> Línea de tarea ("+in.currentTask+") ")
		in.workLineAsTaskLine(line)
	} else if endBracketRegex.MatchString(line) {
		// fmt.Println(in.currentLine, " >> Cierra tarea ("+in.currentTask+") ")
		in.WorkLineCloseTask()
	} else {
		// es otra cosa
		return
	}

}

func (in *Interpreter) workLineAsRootpath(line string) {

	rootPathRegex := regexp.MustCompile(`^\s*\$ROOTPATH\s*=\s*"([^"]+)"`)
	if matches := rootPathRegex.FindStringSubmatch(line); len(matches) == 2 {
		// matches[1] contiene el valor encontrado entre las comillas
		in.Script_rootpath = strings.TrimSpace(matches[1])
	}

}

func (in *Interpreter) workLineAsVariable(line string) {

	if in.Script_rootpath == "" {
		in.workLineAsRootpath(line)
	}

	// Define una expresión regular para encontrar la línea que contiene "$binary_name = "app_api""
	variableRegex := regexp.MustCompile(`^\s*\$[a-zA-Z_]+\s*=\s*"([^"]+)"`)
	matches := variableRegex.FindStringSubmatch(line)
	if len(matches) == 2 {
		// matches[1] contiene el valor encontrado entre las comillas
		in.Script_variables[in.getVariableName(line)] = strings.TrimSpace(matches[1])
	}
}

func (in *Interpreter) getVariableName(line string) string {
	return strings.TrimSpace(strings.Split(line, "=")[0])
}

func (in *Interpreter) workLineAsTask(line string) {

	// Si la línea contiene el patrón, agrégala al slice
	taskName := strings.Split(line, "{")[0]
	in.Script_tasks = append(in.Script_tasks, taskName)
	in.workingCurrentTask = strings.ToLower(taskName)

	// println("-------------")
}

func (in *Interpreter) workLineAsTaskLine(line string) {
	// in.Script_variables[in.currentTask] = os.ModeAppend.String()
	dollarRegex := regexp.MustCompile(`\$`)
	// si la línea contiene una variable entonces reemplazo el nombre de la variable por el valor
	if dollarRegex.MatchString(line) {
		line = in.replaceVariableNameForValue(line)
	}
	in.Script_taskLines[in.workingCurrentTask] = append(in.Script_taskLines[in.workingCurrentTask], strings.TrimSpace(line))
}

func (in *Interpreter) replaceVariableNameForValue(line string) string {
	// comprueba que la variable exista
	// line_aux := strings.Split(line, "$")
	varName := "$" + in.getVariableName(strings.Split(line, "$")[1])
	// println(varName)
	if _, exists := in.Script_variables[varName]; exists {
		return strings.ReplaceAll(line, varName, in.Script_variables[varName])
	}
	return line
}

func (in *Interpreter) WorkLineCloseTask() {
	in.isWorkingOnTask = false
	in.workingCurrentTask = ""
}

func (in *Interpreter) CheckIfTaskExists(taskName string) error {
	if exists := slices.Contains(in.Script_tasks, taskName); exists {
		in.task2run = strings.ToLower(taskName)
		return nil
	}
	err := fmt.Errorf("[error] task \"%s\" does not exists", taskName)
	return err
}

func (in *Interpreter) GetTaskLines(taksName string) []string {
	return in.Script_taskLines[strings.ToLower(taskName)]

}

func (in *Interpreter) RunTask() {
	commandRunner := newCommandRunner(in.Script_rootpath)
	taskLines := in.GetTaskLines(in.task2run)
	commandRunner.RunLines(taskLines)
	in.task2run = ""
}

func (in *Interpreter) Test() {
	// Imprime el mapa de nombres binarios
	fmt.Println("Script ROOTPATH: \"" + in.Script_rootpath + "\"")
	fmt.Println("----------")
	// Imprime el mapa de nombres binarios
	fmt.Println("Nombres de variables encontrados:")
	for key, value := range in.Script_variables {
		fmt.Printf("%s: %s\n", key, value)
	}
	fmt.Println("----------")
	// Imprime el mapa de nombres binarios
	fmt.Println("Nombres de tareas encontrados:")
	for key, value := range in.Script_tasks {
		fmt.Printf("%v: %s\n", key, value)
	}

	fmt.Println("----------")
	// Imprime el mapa de nombres binarios
	fmt.Println("Líneas de tareas encontrados:")
	for key, value := range in.Script_taskLines {
		fmt.Printf("(%v) %s\n", key, value)
	}

	fmt.Println("Total de líneas :", in.currentLine)
}
