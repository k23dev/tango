/*
Copyright © 2024 NAME HERE elanticrypt0@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"tango_cli/filemaker"
	"tango_cli/parser"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tango_cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creación de archivos individuales de features, models y routes",
	Long:  `Crear features, models, views, Api`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
		}
	},
}

var createPackCmd = &cobra.Command{
	Use:   "createPack",
	Short: "Creación de archivos paquetes de features, models y routes",
	Long:  `Crear paquetes features, models, views, Api`,
	Run: func(cmd *cobra.Command, args []string) {

		p := parser.New()
		var fm *filemaker.FileMaker
		var packageName string
		var templateSelected string

		if len(args) > 0 {
			packageName = args[0]
		}

		if len(args) > 1 {
			templateSelected = args[1]
		}

		p.Read(packageName)
		fm = filemaker.New(packageName)

		// PATH
		fm.SetRootPath("./api/")
		fm.SetAppDir("app")
		// aca se define que se crea
		fm.SetMode(templateSelected)
		// modo forzado
		fm.SetForcedMode(true)

		// Creación
		fmt.Println("Making: ", os.Args[1])
		fmt.Println("Mode: ", os.Args[2])
		fmt.Println("Execuit it!")
		fm.MakeIt()
	},
}

var createPackApiCmd = &cobra.Command{
	Use:   "createPackApi",
	Short: "Creación de archivos paquetes de features, models y routes",
	Long:  `Crear paquetes features, models, views, Api`,
	Run: func(cmd *cobra.Command, args []string) {

		p := parser.New()
		var fm *filemaker.FileMaker
		var packageName string
		var templateSelected string = "api"

		if len(args) > 0 {
			packageName = args[0]
		}

		// if len(args) > 1 {
		// templateSelected = args[1]
		// }

		p.Read(packageName)
		fm = filemaker.New(packageName)

		// PATH
		fm.SetRootPath("./api/")
		fm.SetAppDir("app")
		// aca se define que se crea
		fm.SetMode(templateSelected)
		// modo forzado
		fm.SetForcedMode(true)

		// Creación
		fmt.Println("Making: ", os.Args[1])
		fmt.Println("Mode: ", os.Args[2])
		fmt.Println("Execuit it!")
		fm.MakeIt()

	},
}

var createModelCmd = &cobra.Command{
	Use:   "createModel",
	Short: "Crear modelos",
	Long:  `Crear modelos`,
	Run: func(cmd *cobra.Command, args []string) {

		p := parser.New()
		var fm *filemaker.FileMaker
		var packageName string
		var templateSelected string = "model"

		if len(args) > 0 {
			packageName = args[0]
		}

		// if len(args) > 1 {
		// templateSelected = args[1]
		// }

		p.Read(packageName)
		fm = filemaker.New(packageName)

		// PATH
		fm.SetRootPath("./api/")
		fm.SetAppDir("app")
		// aca se define que se crea
		fm.SetMode(templateSelected)
		// modo forzado
		fm.SetForcedMode(true)

		// Creación
		fmt.Println("Making: ", os.Args[1])
		fmt.Println("Mode: ", os.Args[2])
		fmt.Println("Execuit it!")
		fm.MakeIt()

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	appBanner("0.9.1")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tango_cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(createPackCmd)
	rootCmd.AddCommand(createPackApiCmd)
	rootCmd.AddCommand(createModelCmd)
}
