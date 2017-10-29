package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const Version = "0.0.2"

func ShowVersion() {
	fmt.Println("ssm-edit v" + Version)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "ssm-edit version",
	Long:  `ssm-edit version`,
	Run: func(cmd *cobra.Command, args []string) {
		ShowVersion()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
