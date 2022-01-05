/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gh-labeler/pkg/core"

	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync command syncs the repository labels from input label file",
	Long:  `sync command syncs the repository labels from input label file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := core.SyncLabels(cmd, args); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.Flags().StringP("inputFilePath",
		"i", "",
		"absolute path for the output file")
	syncCmd.MarkFlagRequired("inputFilePath")
}
