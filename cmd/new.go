/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: cmdnew,
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func cmdnew(cmd *cobra.Command, args []string) {
    targetPath := getProjectPath(args)
    makeDir(targetPath)
    createCpp(targetPath)
    fmt.Printf("Created at %s", targetPath)
}

func getProjectPath(args []string) string {
    dir, err := os.Getwd()
    if err  != nil {
	log.Fatal(err)
    }

    if len(args) >= 1 {
	dir += "/" + args[0]
    }

    return dir
}

func makeDir(path string) {
    os.Mkdir(path, 0751)
}

func createCpp(path string) {
    initialContent := `#include <iostream>

using namespace std;

int main() {
    ios::sync_with_stdio(0);
    cin.tie(0);
}`
    filename := path + "/" + "main.cpp"

    os.WriteFile(filename, []byte(initialContent), 0644)
}
