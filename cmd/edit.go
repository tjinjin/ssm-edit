package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	edit "github.com/tjinjin/ssm-edit/cli"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit SSM parameter",
	Long:  `Edit SSM parameter`,
	Run: func(cmd *cobra.Command, args []string) {
		p := viper.GetString("profile")
		r := viper.GetString("region")
		n := viper.GetString("name")

		if len(n) == 0 {
			fmt.Println("--name flag required")
			os.Exit(1)
		}
		edit.Edit(p, r, n)
	},
}

func init() {
	RootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.
	editCmd.Flags().StringP("name", "n", "", "Specify name [required]")
	viper.BindPFlag("name", editCmd.Flags().Lookup("name"))

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
