package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "minihuna",
	Short: "minihuna is otasukeman for hunachi.",
	Long: `TODO: describe detail of 'minihuna'.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO:
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

func AddToRootCmd(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func init(){
	// どんなコマンドあるのかを表示してくれる．
	// rootCmd.PersistentFlags().StringP("version", "v", "", "show version")
}
