/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gh-labeler/pkg/core"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan fetches all the labels from a repository and persists in a file",
	Long:  `scan fetches all the labels from a repository and persists in a file`,
	Run: func(cmd *cobra.Command, args []string) {
		err := core.ScanRepoForLabels(cmd, args)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringP("outputFilePath",
		"o", "",
		"absolute path for the output file")
}
