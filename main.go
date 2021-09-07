package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var test int = 1

func main() {

	var cmdBuild = &cobra.Command{
		Use:   "build",
		Short: "Build anything to the screen",
		Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				build("./app/main.go", "app/app")
				build("./auth/main.go", "auth/auth")
			}
		},
	}

	var cmdRun = &cobra.Command{
		Use:   "run",
		Short: "Run anything to the screen",
		Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				go run("./app/app")
				go run("./auth/auth")
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdBuild, cmdRun)
	rootCmd.Execute()
}

func build(path, fileName string) error {
	cmdi := exec.Command("go", "build", "-o "+fileName, path)
	cmdi.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmdi.Stdout = &out
	err := cmdi.Run()
	if err != nil {
		return err
	}
	fmt.Printf("in all caps: %q\n", out.String())
	return nil
}

func run(path string) error {
	cmdi := exec.Command(path)
	cmdi.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmdi.Stdout = &out
	err := cmdi.Run()
	if err != nil {
		return err
	}
	fmt.Printf("in all caps: %q\n", out.String())
	return nil
}
