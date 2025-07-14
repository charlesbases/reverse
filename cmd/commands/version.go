package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v1.0.0"

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print reverse version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reverse version", version)
	},
}
