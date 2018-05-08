package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var profile string

var region string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ssm-edit",
	Short: "Edit SSM parameter store on your EDITOR",
	Long:  `Edit SSM parameter store on your EDITOR`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ssm-edit.yaml)")
	RootCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "default", "~/.aws/credentials (default is default)")
	RootCmd.PersistentFlags().StringVarP(&region, "region", "r", "ap-northeast-1", "AWS REGION(default is ap-northeast-1)")
	viper.BindPFlag("profile", RootCmd.PersistentFlags().Lookup("profile"))
	viper.BindPFlag("region", RootCmd.PersistentFlags().Lookup("region"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
