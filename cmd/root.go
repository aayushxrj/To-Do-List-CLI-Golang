/*
Copyright © 2024 Aayush Ranjan
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pluto",
	Short: "Pluto ia a todo application",
	Long: `Pluto will help you get more done in less time.
It's designed to be as simple as possible to help you accomplish your goals.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var dataFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile string

func initConfig(){
	viper.SetConfigFile(".pluto")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	viper.SetEnvPrefix("pluto")

	//If a config file is found. read it in
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}


func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pluto.yaml)")

	//added
	home, err := homedir.Dir()
	if err != nil{
		log.Println("Unable to detect home directory. Please set data file using --datafile")
	}
	rootCmd.PersistentFlags().StringVar(&dataFile, 
		"datafile",
		home+string(os.PathSeparator)+".tasks.json",
		"data file to store todos")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}