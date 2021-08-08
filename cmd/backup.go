package cmd

import (
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:     "backup",
	Aliases: []string{"bup"},
	Short:   "Backup file",
	Long:    "TODO",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	backupCmd.AddCommand(backupSaveCmd)
	backupCmd.AddCommand(backupRestoreCmd)
}