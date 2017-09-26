// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
