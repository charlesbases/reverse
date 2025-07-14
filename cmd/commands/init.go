package commands

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/charlesbases/reverse/internal/infras/config"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Init reverse yaml",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("file")
		file, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			return err
		}

		file.WriteString(config.DefaultDBConfig)
		file.WriteString(config.DefaultTargetConfig)
		return file.Close()
	},
}
