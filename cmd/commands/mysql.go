package commands

import (
	"github.com/spf13/cobra"
)

var mysqlCommand = &cobra.Command{
	Use:   "mysql",
	Short: "MySQL driver",
	RunE: func(cmd *cobra.Command, args []string) error {
		file, _ := cmd.Flags().GetString("file")
		return reverse(cmd.Context(), file, "mysql")
	},
}
