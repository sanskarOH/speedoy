/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/sanskarOH/nutwrk/internal/test"

	"github.com/spf13/cobra"
)

var size string

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "to test the download and upload speed",
	Long: `It is used to test the download and upload speed of the internet
			Usage = nutwrk test`,
	Run: func(cmd *cobra.Command, args []string) {
		test.Check(size)

	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(
		&size,
		"size",
		"s",
		"100mb",
		"size of the test file downloaded",
	)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
