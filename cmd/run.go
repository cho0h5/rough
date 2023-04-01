/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"bytes"
	"os/exec"
	"fmt"
	"strings"
	"log"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: cmdrun,
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func cmdrun(cmd *cobra.Command, args []string) {
    workdir := getProjectPath(args)
    compileAndRun(workdir)
}

func compileAndRun(workdir string) {
    cmd := exec.Command("g++", "main.cpp", "-std=c++20", "-o", "main")
    var outb, errb bytes.Buffer
    cmd.Stdout = &outb
    cmd.Stderr = &errb
    cmd.Dir = workdir
    err := cmd.Run()

    fmt.Print(strings.Trim(outb.String(), "\n"))
    fmt.Print(strings.Trim(errb.String(), "\n"))
    if err != nil {
	log.Fatal(err)
    }


    cmd = exec.Command("./main")
    cmd.Stdout = &outb
    cmd.Stderr = &errb
    cmd.Stdin = os.Stdin
    cmd.Dir = workdir
    err = cmd.Run()

    fmt.Print(strings.Trim(outb.String(), "\n"))
    fmt.Print(strings.Trim(errb.String(), "\n"))
    if err != nil {
	log.Fatal(err)
    }
}
