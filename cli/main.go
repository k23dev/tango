package main

import (
	"fmt"
	"os"
	"tango_cli/filemaker"
)

func main() {

	appBanner()

	paramsSetted := 0

	//Example: tango_cli photo basic f

	var fm *filemaker.FileMaker

	if len(os.Args) > 1 {

		fm = filemaker.New(os.Args[1])
		fmt.Println(os.Args[0])
		fm.SetRootPath("../api")
		fm.SetAppDir("app")

		paramsSetted = paramsSetted + 1

	}

	if len(os.Args) >= 2 {

		// aca define que va a crear
		fm.SetMode(os.Args[2])
		paramsSetted = paramsSetted + 1
	}

	if len(os.Args) >= 4 {

		// aca define el modo SI es forzado o no
		fmt.Println("Forced Mode")
		fm.SetForcedMode(true)

		paramsSetted = paramsSetted + 1
	}

	// aca ejecuta todo

	if paramsSetted >= 2 {

		fmt.Println("Making: ", os.Args[1])
		fmt.Println("Mode: ", os.Args[2])
		fmt.Println("Making files...")
		fm.MakeIt()
		fmt.Println("-----------------")
		fmt.Println("- Add the routes call to the app/routes/setupapproutes.go")
		fmt.Println("- Add the name on the main menu at app/views/nav.temp")
	}

}

func appBanner() {

	fmt.Println(" ### ### ### ####")
	fmt.Println(" ### ### ### ####")
	fmt.Println("    TANGO_CLI    ")
	fmt.Println(" ### ### ### ####")
	fmt.Println(" ### ### ### ####")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("- basic: Route, Feature, Model")
	fmt.Println("- full: Route, Feature, Model, View")
	fmt.Println("- fullWithselector: Route, Feature, Model, View, Selector")
	fmt.Println("")
	fmt.Println("e.g. [NAME] [MODE] [(optional)FORCED_MODE]")

}
