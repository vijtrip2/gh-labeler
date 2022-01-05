/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gh-labeler",
	Short: "gh-labeler helps in managing the GitHub labels for your repository",
	Long:  `gh-labeler helps in managing the GitHub labels for your repository`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("I am gh-labeler")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("repo", "", "github repo name")
	rootCmd.MarkPersistentFlagRequired("repo")
	rootCmd.PersistentFlags().String("org", "", "github org name")
	rootCmd.MarkPersistentFlagRequired("org")
	rootCmd.PersistentFlags().Bool("dryrun", false, "")
}
