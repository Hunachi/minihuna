package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/user"
)

func init() {
	// rootCmd.AddCommand(saveCmd)
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save setting",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		saveSetting()
	},
}

func saveSetting() {
	if err := writeByres("settings.txt", "localhost 10.10.10.10 ~/.ssh/id_rsa2\n"); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func writeByres(filename string, setting string) error {
	current, err := user.Current()
	if err != nil {
		fmt.Println("Error: Where are you? (user.Current is nil !!)")
		return err
	}

	path := current.HomeDir + "/.minihuna/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path , 0755)
	}
	if err != nil {
		return err
	}

	file, err := os.Create(path + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	lines = append(lines, setting)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("EOF", file.Name())
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("Reading...", line)
		lines = append(lines, line)
	}

	for _, line := range lines {
		fmt.Println(line)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
