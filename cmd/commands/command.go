package commands

import (
	"log"
	"os"
)

// Init adds one or more commands to this parent command.
func Init() () {
	// add flags
	rootCommand.PersistentFlags().StringP("file", "f", "./reverse.yaml", "reverse file")

	// add subcommands
	rootCommand.AddCommand(versionCommand)
	rootCommand.AddCommand(initCommand)
	rootCommand.AddCommand(mysqlCommand)
}

// Execute command run
func Execute() () {
	if err := rootCommand.Execute(); err != nil {
		log.Println(err)
		os.Exit(0)
	}
}
