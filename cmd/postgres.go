package cmd

import (
	"github.com/spf13/cobra"
)

var postgresCommand = &cobra.Command{
	Use:   "postgres",
	Short: "PostgreSQL driver",
	RunE: func(cmd *cobra.Command, args []string) error {
		file, _ := cmd.Flags().GetString("file")
		return reverse(cmd.Context(), file, "postgres")
	},
}
