package cmd

import (
	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use: "show",
	Short: "Show minihuna's features.",
	Long: "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		showAllFeatures()
	},
}

func showAllFeatures(){
	println("======== minihuna's Features ========")
	for _, command := range rootCmd.Commands() {
		println(command.Name())
	}
	println("======== Finish ========")
}
