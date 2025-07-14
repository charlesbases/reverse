package commands

import (
	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "reverse",
	Short: "Reverse",
	Long:  "Generates a golang struct for the specified schema and table.",
}
