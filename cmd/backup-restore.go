package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/user"
)

var backupRestoreCmd = &cobra.Command{
	Use:     "restore",
	Aliases: []string{"rst"},
	Short:   "Restore file",
	Long:    "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" {
			fmt.Println("ファイルを指定してください")
			return
		}
		if len(args) <= 1 || args[1] == "" {
			args = append(args, args[0])
		}
		srcName := args[0]
		dstName := args[1]
		if dstName == "" {
			dstName = srcName
		}
		err := restoreFile(srcName, dstName)

		if err != nil {
			panic(err)
		}
		println("リストアが完了したよ！", args[0], " -> ", args[1])
	},
}

func restoreFile(srcName string, dstName string) error {

	current, err := user.Current()
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	path := current.HomeDir + "/.minihuna/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
	}
	if err != nil {
		return err
	}

	srcPath := current.HomeDir + "/.minihuna/" + srcName
	dstPath := dir + "/" + dstName

	return copyFile(srcPath, dstPath)
}
