package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	list "github.com/tjinjin/ssm-edit/cli"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list ssm parameters",
	Long:  `list ssm parameters`,
	Run: func(cmd *cobra.Command, args []string) {
		p := viper.GetString("profile")
		r := viper.GetString("region")
		list.List(p, r)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
