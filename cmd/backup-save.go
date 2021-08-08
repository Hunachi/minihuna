package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/user"
)

var backupSaveCmd = &cobra.Command{
	Use:     "save",
	Short:   "Backup file",
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
		err := backupFile(srcName, dstName)

		if err != nil {
			panic(err)
		}
		println("バックアップが完了したよ！", args[0], " -> ", args[1])
	},
}

func copyFile(srcPath string, dstPath string) error {

	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func backupFile(srcName string, dstName string) error {

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

	srcPath := dir + "/" + srcName
	dstPath := current.HomeDir + "/.minihuna/" + dstName

	return copyFile(srcPath, dstPath)
}
